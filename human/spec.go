package main

import (
	"gopkg.in/yaml.v2"
)

// Command is a runnable command
type Command struct {
	Exe  string
	Args []string
}

// TestCase represents a single test case in a spec
type TestCase struct {
	It     string
	Run    Command
	Expect yaml.MapSlice
}

// Spec represents a test suite
type Spec struct {
	Scenario string
	Tests    []TestCase
}
