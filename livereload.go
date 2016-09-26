package main

import (
	"html/template"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/bluesuncorp/wash/env"
	"github.com/go-playground/log"
	"github.com/go-playground/statics/static"

	"github.com/jaschaephraim/lrserver"
	"gopkg.in/fsnotify.v1"
)

// startLiveReloadServer initializes a livereload to notify the browser of changes to code that does not need a recompile.
func startLiveReloadServer(tpls *template.Template, cfg *env.Config, staticAssets *static.Files) {

	if cfg.IsProduction {
		return
	}

	log.Info("Initializing livereload")

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalf("%s %s%s%s", UnknownError, Red, err.Error(), Reset)
	}

	defer watcher.Close()

	walker := func(path string, info os.FileInfo, err error) error {

		if info.IsDir() {
			err = watcher.Add(path)
			if err != nil {
				return err
			}
		}

		return nil
	}

	err = filepath.Walk("assets", walker)
	if err != nil {
		log.WithFields(log.F("error", err)).Fatal("Failed to Walk assets for livereload")
	}

	err = filepath.Walk("templates", walker)
	if err != nil {
		log.WithFields(log.F("error", err)).Fatal("Failed to Walk assets for livereload")
	}

	done := make(chan bool)

	lr, err := lrserver.New(lrserver.DefaultName, lrserver.DefaultPort)
	if err != nil {
		log.Error(err)
	}

	// Start LiveReload server
	go func() {
		err := lr.ListenAndServe()
		if err != nil {
			log.WithFields(log.F("error", err)).Error("error with livereload")
		}
	}()

	var locker sync.Mutex
	timerRunning := false

	eventMap := map[string]*fsnotify.Event{}

	go func() {
		for {
			select {
			case event := <-watcher.Events:

				locker.Lock()
				eventMap[event.Name] = &event

				if !timerRunning {
					timerRunning = true

					go func() {

						time.Sleep(200 * time.Millisecond)

						locker.Lock()

						for _, event := range eventMap {

							ext := filepath.Ext(event.Name)

							if ext == ".js" {

								log.Infof("%s %sJavascript Updated: %s%s\n", ThumbsUpEmoji, Green, event.Name, Reset)

								lr.Reload(event.Name)
								time.Sleep(100 * time.Millisecond)

							} else if ext == ".css" {

								log.Infof("%s %sCSS Updated: %s%s\n", ThumbsUpEmoji, Green, event.Name, Reset)

								lr.Reload(event.Name)
								time.Sleep(100 * time.Millisecond)

							} else if ext == ".tmpl" {

								log.Infof("Compiling Templates: %s\n", event.Name)
								templates, err := initTemplates(cfg, staticAssets)

								if err != nil {
									log.Errorf("%s %sError Compiling Templates: %s%s\n", UnknownError, Red, err.Error(), Reset)
								} else {
									*tpls = *templates

									log.Infof("%s %sTemplates Updated: %s%s\n", ThumbsUpEmoji, Green, event.Name, Reset)
									lr.Reload(event.Name)
									time.Sleep(100 * time.Millisecond)
								}
							}
						}

						eventMap = map[string]*fsnotify.Event{}
						timerRunning = false
						locker.Unlock()
					}()
				}

				locker.Unlock()

			case err := <-watcher.Errors:

				log.Errorf("%s %sWatcher Error:%s%s\n", UnknownError, Red, err.Error(), Reset)
			}
		}
	}()

	<-done
}
