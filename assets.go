//go:generate statics -i=assets -o=assets.go -pkg=main -group=Assets

package main

import "github.com/go-playground/statics/static"

// newStaticAssets initializes a new *static.Files instance for use
func newStaticAssets(config *static.Config) (*static.Files, error) {

	return static.New(config, &static.DirFile{})
}
