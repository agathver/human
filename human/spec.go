package main

// Expectations are in a test case
type Expectations struct {
	ExitCode int
}

// Command is a runnable command
type Command struct {
	Exe  string
	Args []string
}

// TestCase represents a single test case in a spec
type TestCase struct {
	It     string
	Run    Command
	Expect Expectations
}

// Spec represents a test suite
type Spec struct {
	Scenario string
	Tests    []TestCase
}
