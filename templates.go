package main

import (
	"bytes"
	"errors"
	"html/template"

	"github.com/bluesuncorp/wash/env"
	"github.com/go-playground/assets"
	"github.com/go-playground/log"
	"github.com/go-playground/statics/static"
)

const (
	productionJQuery  = "<script src=\"https://ajax.googleapis.com/ajax/libs/jquery/3.1.0/jquery.min.js\"></script>"
	developmentJQuery = "<script src=\"/assets/js/jquery-3.1.0.min.js\"></script>"
	livereloadScript  = "<script src=\"/assets/js/livereload.js?host=localhost\"></script>"

	leftDelim  = "/*include("
	rightDelim = ")*/"
)

// InitTemplates initializes the app's templates
func initTemplates(cfg *env.Config, staticAssets *static.Files) (*template.Template, error) {

	log.Info("Initializing Templates")

	// setup any template functions here
	globalTemplateFunctions := template.FuncMap{
		"jquery": func() template.HTML {
			if cfg.IsProduction {
				return template.HTML(productionJQuery)
			}

			return template.HTML(developmentJQuery)
		},
		"livereload": func() template.HTML {
			if !cfg.IsProduction {
				return template.HTML(livereloadScript)
			}

			return template.HTML("")
		},
		"multi": func(values ...interface{}) (map[string]interface{}, error) {

			if len(values)%2 != 0 {
				return nil, errors.New("invalid multi call")
			}

			params := make(map[string]interface{}, len(values)/2)

			for i := 0; i < len(values); i += 2 {

				key, ok := values[i].(string)
				if !ok {
					return nil, errors.New("key must be a string")
				}

				params[key] = values[i+1]
			}

			return params, nil
		},
	}

	var err error
	var funcs template.FuncMap

	if cfg.IsProduction {

		b, err := staticAssets.ReadFile("/assets/manifest.txt")
		if err != nil {
			return nil, err
		}

		funcs, err = assets.ProcessManifestFiles(bytes.NewBuffer(b), "assets/", assets.Production, true, leftDelim, rightDelim)
		if err != nil {
			return nil, err
		}

	} else {
		funcs, err = assets.LoadManifestFiles("assets/", assets.Development, true, leftDelim, rightDelim)
		if err != nil {
			return nil, err
		}
	}

	for k, v := range funcs {
		globalTemplateFunctions[k] = v
	}

	tpls, err := newStaticTemplates(&static.Config{
		UseStaticFiles: cfg.IsProduction,
		FallbackToDisk: true,
		AbsPkgPath:     appPath,
	})
	if err != nil {
		return nil, err
	}

	log.Info("Reading Templates")

	files, err := tpls.ReadFiles("/templates", true)
	if err != nil {
		return nil, err
	}

	// glob load templates
	buff := new(bytes.Buffer)

	for _, file := range files {
		_, err = buff.Write(file)
		if err != nil {
			log.WithFields(log.F("error", err)).Warn("Issue writing template to buffer")
		}
	}

	tpl, err := template.New("all").Funcs(globalTemplateFunctions).Parse(buff.String())
	if err != nil {
		return nil, err
	}

	return tpl, nil
}
