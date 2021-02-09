package main

import (
	"fmt"

	"github.com/operator-framework/operator-registry/pkg/lib/bundle"
)

func main() {
	// GenerateFunc(directory, outputDir, packageName, channels, channelDefault string, overwrite bool)
	err := bundle.GenerateFunc("manifests", "bundle", "pname", "chans", "alpha", false)
	if err != nil {
		fmt.Println(err.Error())
	}
}
