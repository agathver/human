package main

import (
	"fmt"
	"os/exec"
	"syscall"
)

// RunTests is used to run the tests
func RunTests(spec Spec) {
	fmt.Printf("Scenario %s\n", spec.Scenario)

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
}

func test(t TestCase) (bool, error) {
	cmd := exec.Command(t.Run.Exe, t.Run.Args...)

	err := cmd.Start()

	if err != nil {
		return false, err
	}

	err = cmd.Wait()

	// if err != nil {
	// 	if cmd.ProcessState.Success
	// }

	ws := cmd.ProcessState.Sys().(syscall.WaitStatus)
	exitCode := ws.ExitStatus()

	return t.Expect.ExitCode == exitCode, nil

}
