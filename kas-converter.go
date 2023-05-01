package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func usage() {
	fmt.Printf("Usage: %v -i [-o]\n", filepath.Base(os.Args[0]))
	flag.PrintDefaults()
	os.Exit(0)
}

func main() {
	var (
		inputManifest       = flag.String("i", "", "input repo manifest")
		outputManifest      = flag.String("o", "", "output kas manifest")
		useOutput      bool = false
	)

	flag.Parse()

	if len(*inputManifest) == 0 {
		usage()
	}

	if len(*outputManifest) > 0 {
		useOutput = true
	}

	parseInput(*inputManifest)
	generateOutput(useOutput, *outputManifest)
}

func parseInput(inputManifest string) {
}

func generateOutput(useOutput bool, outputManifest string) {
}
