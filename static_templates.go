//go:generate statics -i=templates -o=static_templates.go -pkg=main -group=Templates

package main

import "github.com/go-playground/statics/static"

// newStaticTemplates initializes a new *static.Files instance for use
func newStaticTemplates(config *static.Config) (*static.Files, error) {

	return static.New(config, &static.DirFile{})
}
