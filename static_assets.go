//go:generate assets -i=assets -extensions=.js,.css -ld=/*include( -rd=)*/ -rtd=true -o=public
//go:generate statics -i=public -o=static_assets.go -pkg=main -group=Assets -ignore=/\. -prefix=public/

package main

import "github.com/go-playground/statics/static"

// newStaticAssets initializes a new *static.Files instance for use
func newStaticAssets(config *static.Config) (*static.Files, error) {

	return static.New(config, &static.DirFile{})
}
