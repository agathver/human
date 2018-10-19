package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os/exec"
	"strings"
	"syscall"

	"gopkg.in/yaml.v2"
)

// RunTests is used to run the tests
func RunTests(spec Spec) {
	header := fmt.Sprintf("Scenario %s", spec.Scenario)
	fmt.Println(header)
	fmt.Println(strings.Repeat("-", len(header)))

	for i, t := range spec.Tests {
		status, err := test(t)

		if err != nil {
			fmt.Printf("    %d ERROR It %s: %s\n", i+1, t.It, err)
			return
		}

		if status {
			fmt.Printf("    %d PASS It %s\n", i+1, t.It)
		} else {
			fmt.Printf("    %d FAIL It %s\n", i+1, t.It)
		}
	}

	fmt.Println()
}

func test(t TestCase) (bool, error) {
	cmd := exec.Command(t.Run.Exe, t.Run.Args...)

	success := true

	expects := make(map[string]interface{})

	for _, item := range t.Expect {
		expects[item.Key.(string)] = item.Value
	}

	expectedExitCode, hasExitCond := expects["exitcode"]
	outputConds, hasOutputCond := expects["output"]

	err := cmd.Start()

	if err != nil {
		return false, err
	}

	if hasOutputCond {
		outputExpects := make(map[string]string)

		for _, item := range outputConds.(yaml.MapSlice) {
			expects[item.Key.(string)] = item.Value.(string)
		}

		success, err = testOutputs(cmd, outputExpects)

		if err != nil {
			return false, err
		}
	}

	cmd.Wait()

	if success && hasExitCond {
		ws := cmd.ProcessState.Sys().(syscall.WaitStatus)
		success = expectedExitCode.(int) == ws.ExitStatus()
	}

	return success, nil

}

func testOutputs(cmd *exec.Cmd, conditions map[string]string) (bool, error) {
	expectedStdout, hasStdout := conditions["stdout"]
	expectedStderr, hasStderr := conditions["stderr"]

	success := true

	if hasStdout {
		stdout, err := cmd.StdoutPipe()

		if err != nil {
			return false, err
		}

		success, err = testPipeMatches(stdout, expectedStdout)

		if err != nil {
			return false, err
		}
	}

	if success && hasStderr {
		stderr, err := cmd.StderrPipe()

		if err != nil {
			return false, err
		}

		success, err = testPipeMatches(stderr, expectedStderr)

		if err != nil {
			return false, err
		}
	}

	return success, nil
}

func testPipeMatches(pipe io.Reader, content string) (bool, error) {
	pipeContent, err := ioutil.ReadAll(pipe)

	if err != nil {
		return false, err
	}

	return string(pipeContent) == content, nil
}
