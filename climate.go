// Copyright (C) 2015 Miquel Sabaté Solà <mikisabate@gmail.com>
// This file is licensed under the MIT license.
// See the LICENSE file.

// TODO: documentation
package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	// TODO: think of something cleaner?
	compact bool
)

func main() {
	backends := []backend{
		&cover{},
		&gofmt{},
		&errcheck{},
		&lint{},
		&vet{},
	}

	for _, backend := range backends {
		backend.setOptions()
	}
	// TODO: fails horribly if we just want to "climate -failsOnly" because the
	// current CLI is utter bullshit.
	flag.BoolVar(&compact, "compact", false, "The results are given in a compact format.")
	flag.Parse()

	all := flag.NFlag() == 0
	hasErrors, executed := false, false
	for _, backend := range backends {
		if backend.installed() && (all || backend.isSet()) {
			if !backend.run() {
				hasErrors = true
			}
			executed = true
		}
	}

	if !executed {
		fmt.Printf("There were no available backends!\n")
	} else if hasErrors {
		os.Exit(1)
	}
	os.Exit(0)
}
