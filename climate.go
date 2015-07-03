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
	flag.BoolVar(&compact, "compact", false, "The results are given in a compact format.")
	flag.Parse()

	// Decide whether to run all the backends are just some of them.
	nflags := flag.NFlag()
	all := nflags == 0 || (nflags == 1 && compact)

	// And finally execute all you can.
	hasErrors, executed := false, false
	for _, backend := range backends {
		if backend.installed() && (all || backend.isSet()) {
			if !backend.run() {
				hasErrors = true
			}
			executed = true
		}
	}

	if hasErrors {
		os.Exit(1)
	}
	if !compact && !executed {
		fmt.Printf("There were no available backends!\n")
	}
	os.Exit(0)
}
