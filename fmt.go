// Copyright (C) 2015 Miquel Sabaté Solà <mikisabate@gmail.com>
// This file is licensed under the MIT license.
// See the LICENSE file.

package main

import (
	"bytes"
	"flag"
	"os/exec"
)

type gofmt struct {
	set bool
}

func (f *gofmt) setOptions() {
	flag.BoolVar(&f.set, "fmt", false, "Use the fmt tool.")
}

func (f *gofmt) installed() bool {
	return true
}

func (f *gofmt) isSet() bool {
	return f.set
}

func (f *gofmt) run(pack string) bool {
	printBackendStatus("fmt")

	// Unfortunately, gofmt needs either the relative or the absolute path, the
	// import path is not enough.
	if pack != "." {
		pack = packageAbs(pack)
	}

	cmd := exec.Command("go", "fmt", pack)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
		printResult("fmt", err.Error(), errored)
		return false
	}

	// And print out the results.
	results := out.String()
	if results == "" {
		printResult("fmt", "", ok)
		return true
	}
	printResult("fmt", out.String(), failed)
	return false
}
