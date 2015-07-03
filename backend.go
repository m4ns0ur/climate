// Copyright (C) 2015 Miquel Sabaté Solà <mikisabate@gmail.com>
// This file is licensed under the MIT license.
// See the LICENSE file.

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mssola/colors"
)

const (
	ok = iota
	failed
	errored
)

// TODO : name
type backend interface {
	setOptions()
	installed() bool
	isSet() bool
	run() bool
}

func printBackendStatus(name string) {
	if !compact {
		c := colors.Default()
		c.SetMode(colors.Bold)
		fmt.Printf(c.Get(fmt.Sprintf("%v: ", name)))
	}
}

// TODO: factor this out, so we don't have to pass the name...
func printResult(name, output string, status int) {
	if !compact {
		prettifyStatus(status)
	}
	if output != "" {
		if compact {
			for _, v := range strings.Split(output, "\n") {
				if v != "" {
					fmt.Printf("%v:%v\n", name, v)
				}
			}
		} else {
			fmt.Printf("%v", output)
		}
	}
}

func prettifyStatus(status int) {
	var fg colors.Colors
	if status == ok {
		fg = colors.Saved
	} else {
		fg = colors.Red
	}

	c := colors.Color{
		Foreground: fg,
		Background: colors.Saved,
		Mode:       colors.Bold,
	}

	switch status {
	case ok:
		fmt.Printf(c.Get("OK\n"))
	case failed:
		fmt.Printf(c.Get("FAILED\n"))
	case errored:
		fmt.Printf(c.Get("ERROR\n"))
	}
}

func packageExists(importPath string) bool {
	if packageExistsIn("GOPATH", importPath) {
		return true
	}
	return packageExistsIn("GOROOT", importPath)
}

func packageExistsIn(env, importPath string) bool {
	base := os.Getenv(env)
	if base == "" {
		return false
	}

	path := filepath.Join(base, "src", importPath)
	_, err := os.Stat(path)
	return err == nil
}
