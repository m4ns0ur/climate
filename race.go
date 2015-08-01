// Copyright (C) 2015 Miquel Sabaté Solà <mikisabate@gmail.com>
// This file is licensed under the MIT license.
// See the LICENSE file.

package main

import (
	"bytes"
	"flag"
	"os/exec"
)

type race struct {
	set bool
}

func (r *race) setOptions() {
	flag.BoolVar(&r.set, "race", false, "Use the 'go build -race' tool.")
}

func (r *race) installed() bool {
	return true
}

func (r *race) isSet() bool {
	return r.set
}

func (r *race) run(pack string) bool {
	printBackendStatus("race")

	// Unfortunately, go build needs either the relative or the absolute path,
	// the import path is not enough.
	if pack != "." {
		pack = packageAbs(pack)
	}

	cmd := exec.Command("go", "test", "-race", pack)
	var out, errOut bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errOut
	err := cmd.Run()

	if err != nil {
		if str := out.String(); str != "" {
			printResult("race", str, failed)
		} else {
			printResult("race", errOut.String(), failed)
		}
		return false
	}
	printResult("race", "", ok)
	return true
}
