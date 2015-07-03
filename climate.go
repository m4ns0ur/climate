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
