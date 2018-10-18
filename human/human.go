package main

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

func usage() {
	println("human - Pseudo human to test your command line tools")
	println("Version: 0.0.1 - completely untested code")
	println()
	println("Usage: human SPEC")
	println()
	println("Arguments:")
	println("    SPEC: Spec file with tests")
}

func main() {

	if len(os.Args) == 1 {
		usage()
	}

	specFile := os.Args[1]

	input, err := ioutil.ReadFile(specFile)

	if err != nil {
		panic(err)
	}

	var specs []Spec

	err = yaml.Unmarshal(input, &specs)

	if err != nil {
		panic(err)
	}

	for _, spec := range specs {
		RunTests(spec)
	}
}
