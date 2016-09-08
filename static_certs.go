//go:generate statics -i=certs -o=certs.go -pkg=main -group=Certs -ignore=/\. -prefix=certs/

package main

import "github.com/go-playground/statics/static"

// newStaticCerts initializes a new *static.Files instance for use
func newStaticCerts(config *static.Config) (*static.Files, error) {

	return static.New(config, &static.DirFile{})
}
