// Copyright (C) 2015 Miquel Sabaté Solà <mikisabate@gmail.com>
// This file is licensed under the MIT license.
// See the LICENSE file.

package main

import (
	"bytes"
	"flag"
	"os/exec"
)

type vet struct {
	set bool
}

func (v *vet) setOptions() {
	flag.BoolVar(&v.set, "vet", false, "Use the vet tool.")
}

func (v *vet) installed() bool {
	return packageExists("golang.org/x/tools/cmd/vet")
}

func (v *vet) isSet() bool {
	return v.set
}

func (v *vet) run() bool {
	printBackendStatus("vet")

	cmd := exec.Command("go", "vet")
	var out bytes.Buffer
	cmd.Stderr = &out
	err := cmd.Run()

	if err != nil {
		printResult(out.String(), failed)
		return false
	}
	printResult("", ok)
	return true
}
