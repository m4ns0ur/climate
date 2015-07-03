// Copyright (C) 2015 Miquel Sabaté Solà <mikisabate@gmail.com>
// This file is licensed under the MIT license.
// See the LICENSE file.

package main

import (
	"bytes"
	"flag"
	"os/exec"
)

type lint struct {
	set bool
}

func (l *lint) setOptions() {
	flag.BoolVar(&l.set, "lint", false, "Use the golint tool.")
}

func (l *lint) installed() bool {
	return packageExists("github.com/golang/lint/golint")
}

func (l *lint) isSet() bool {
	return l.set
}

func (l *lint) run() bool {
	printBackendStatus("lint")

	cmd := exec.Command("golint", ".")
	var out bytes.Buffer
	cmd.Stdout = &out
	_ = cmd.Run()

	if out.String() != "" {
		printResult("lint", out.String(), failed)
		return false
	}
	printResult("lint", "", ok)
	return true
}
