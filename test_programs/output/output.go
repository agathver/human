package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	fmt.Println(os.Args[1])

	if len(os.Args) == 3 {
		code, _ := strconv.ParseInt(os.Args[2], 0, 32)
		os.Exit(int(code))
	}

	// exit with 0
}
