package main

import (
	"fmt"

	"github.com/operator-framework/operator-registry/pkg/lib/bundle"
)

func main() {
	// GenerateFunc(directory, outputDir, packageName, channels, channelDefault string, overwrite bool)
	err := bundle.GenerateFunc("manifests/0.0.1", "bundle/bundle-0.0.1", "pname", "chans", "alpha", false)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = bundle.GenerateFunc("manifests/0.0.2", "bundle/bundle-0.0.2", "pname", "chans", "alpha", false)
	if err != nil {
		fmt.Println(err.Error())
	}
}
