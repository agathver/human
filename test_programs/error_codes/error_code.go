package main

import (
	"os"
	"strconv"
)

func main() {

	if len(os.Args) == 1 {
		os.Exit(0)
	}

	code, error := strconv.ParseInt(os.Args[1], 0, 32)

	if error != nil {
		os.Exit(255)
	}

	os.Exit(int(code))
}
