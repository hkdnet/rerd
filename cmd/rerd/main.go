package main

import (
	"fmt"
	"os"

	"github.com/hkdnet/rerd"
)

const Usage = "rerd FILENAME"

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, Usage)
		os.Exit(1)
	}
	err := rerd.Run(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
