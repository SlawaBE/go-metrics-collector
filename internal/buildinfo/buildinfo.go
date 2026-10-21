// Package buildinfo provides a method Print for writing information about
// the version, date, and commit of the build to stdout.
package buildinfo

import "fmt"

// Print writes the build version, date and commit to stdout, using "N/A"
// for any value that was not set at build time.
func Print(version string, date string, commit string) {
	fmt.Printf("Build version: %s\n", orNA(version))
	fmt.Printf("Build date: %s\n", orNA(date))
	fmt.Printf("Build commit: %s\n", orNA(commit))
}

func orNA(value string) string {
	if value == "" {
		return "N/A"
	}
	return value
}
