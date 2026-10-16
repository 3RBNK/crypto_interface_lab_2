package main

import (
	"filecrypt/internal/cli"
	"fmt"
	"os"
)

func main() {
	if err := cli.Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintln(os.Stderr,
			"Usage: filecrypt <encrypt|decrypt> --in <file> --out <file> [--no-verify]")
		os.Exit(1)
	}
}
