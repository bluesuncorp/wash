package main

import (
	"html/template"

	"github.com/bluesuncorp/wash/env"
	"github.com/go-playground/livereload"
	"github.com/go-playground/log"
	"github.com/go-playground/statics/static"
)

// startLiveReloadServer initializes a livereload to notify the browser of changes to code that does not need a recompile.
func startLiveReloadServer(tpls *template.Template, cfg *env.Config, staticAssets *static.Files) error {

	if cfg.IsProduction {
		return nil
	}

	log.Info("Initializing livereload")

	paths := []string{
		"assets",
		"templates",
	}

	tmplFn := func(name string) (bool, error) {

		templates, err := initTemplates(cfg, staticAssets)
		if err != nil {
			return false, err
		}

		*tpls = *templates

		return true, nil

	}

	mappings := livereload.ReloadMapping{
		".css":  nil,
		".js":   nil,
		".tmpl": tmplFn,
	}

	_, err := livereload.ListenAndServe(livereload.DefaultPort, paths, mappings)

	return err
}
