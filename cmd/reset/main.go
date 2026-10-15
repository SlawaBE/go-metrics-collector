// Package main implements code generation of Reset() methods for structs
// marked with the // generate:reset directive.
//
// # What it does
//
// The utility walks the project directories starting from the root and scans
// all Go packages below it. For every struct carrying the // generate:reset
// comment it generates a Reset() method that returns the object to its zero
// (initial) state:
//
//   - primitive fields are reset to their zero value (int to 0, string to "",
//     bool to false, and so on);
//   - slices are truncated by length, not nilled (strings to strings[:0]);
//   - maps are cleared with the built-in clear function;
//   - non-nil pointers are reset according to the same rules;
//   - nested structs declared in the same package call Reset() when they have
//     one or are themselves marked with // generate:reset, otherwise the
//     generator fails with an error pointing to the offending field.
//
// The generated methods for the structs of one package are written into a
// single reset.gen.go file in that package. The file is rewritten completely
// on every run.
//
// # How to run
//
// Build the binary from the cmd/reset directory:
//
//	go build -o reset ./cmd/reset
//
// Then run it from the project root:
//
//	./reset
//
// Optionally pass a root directory as the first argument to scan a
// subdirectory instead:
//
//	./reset ./internal/model
//
// # Limitations
//
// The generator targets data-only structs. Channels, functions, interfaces
// and external (imported) types are not supported; encountering such a field
// aborts generation with an error indicating its position.
package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

// run walks the directory tree from the given root (default ".") and processes
// each package directory.
func run() error {
	root := "."
	if len(os.Args) > 1 {
		root = os.Args[1]
	}

	if info, err := os.Stat(root); err != nil {
		return fmt.Errorf("cannot access root %s: %w", root, err)
	} else if !info.IsDir() {
		return fmt.Errorf("root %s is not a directory", root)
	}

	return filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			return nil
		}
		if path != root && skipScanDir(d.Name()) {
			return filepath.SkipDir
		}
		if err := processDir(path); err != nil {
			return fmt.Errorf("reset generation: %w", err)
		}
		return nil
	})
}
