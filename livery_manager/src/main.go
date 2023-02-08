package main

import (
	"fmt"
)

func main() {
	clearErr := removeOutputDir()
	if(clearErr != nil) {
		fmt.Printf("Error clearing output directory: %s", clearErr)
	}

	config := ReadConfigFromJson("./files/config.json")
	err := generate(config)
	if(err!=nil) {
		fmt.Printf("%s \n", err.Error())
	}
}
