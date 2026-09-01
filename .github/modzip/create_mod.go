// Command create_mod produces a Go module source zip byte-compatible with the
// one proxy.golang.org serves, using the proxy's own packing implementation.
//
// It lives under .github/ so that `go build ./...` at the module root ignores
// it (the go tool skips directories beginning with a dot) and so that the
// package step's `rm -rf .github` keeps it out of the module zip.
package main

import (
	"log"
	"os"

	"golang.org/x/mod/module"
	"golang.org/x/mod/zip"
)

func main() {
	if len(os.Args) != 5 {
		log.Fatal("usage: create_mod <module-path> <version> <source-dir> <output-zip>")
	}
	m := module.Version{Path: os.Args[1], Version: os.Args[2]}
	f, err := os.Create(os.Args[4])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := zip.CreateFromDir(f, m, os.Args[3]); err != nil {
		log.Fatal(err)
	}
	log.Printf("created module zip: %s", os.Args[4])
}
