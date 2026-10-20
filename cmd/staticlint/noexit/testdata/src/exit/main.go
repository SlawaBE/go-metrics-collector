// Package main contains test code for the noexit analyzer. Files under
// testdata directories are compiled by analysistest and are not part of the
// main project build.
package main

import (
	"fmt"
	"os"
)

func exitFromOtherFunction() {
	// Calling os.Exit outside the main function is not a violation.
	os.Exit(1)
}

func main() {
	fmt.Println("start")

	os.Exit(0) // want "direct call to os.Exit in main"
}
