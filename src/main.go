package main

import (
	"flag"
	"fmt"
	"./config"
	"./generator"
)

func main() {
	cleanPtr := flag.Bool("clean", false, "clear output directory before generating the packages")
	versionPtr := flag.String("version", "DAR_1.0", "version string to be set (should be used to prevent collisions)")

	if(*cleanPtr) {
		clearErr := generator.RemoveOutputDir()
		fmt.Printf("Error clearing output directory: %s", clearErr)
	}

	demoConfig := config.ReadConfigFromJson("./config.json")
	generator.Generate(demoConfig, *versionPtr)
}
