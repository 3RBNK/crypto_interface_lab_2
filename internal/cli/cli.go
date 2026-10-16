package cli

import (
	"flag"
	"fmt"
)

type CryptoMode int

const (
	Encrypt CryptoMode = iota
	Decrypt
)

func (m CryptoMode) String() string {
	switch m {
	case Encrypt:
		return "encrypt"
	case Decrypt:
		return "decrypt"
	default:
		return "unknown"
	}
}

type InputContext struct {
	inputFilePath  string
	outputFilePath string
	cryptoMode     CryptoMode
	noVerify       bool
}

func readTermParams(args []string) (InputContext, error) {
	var ctx InputContext

	if len(args) == 0 {
		return ctx, fmt.Errorf("crypto mode not specified")
	}

	switch args[0] {
	case "encrypt":
		ctx.cryptoMode = Encrypt
	case "decrypt":
		ctx.cryptoMode = Decrypt
	default:
		return ctx, fmt.Errorf("crypto mode not specified: %s", args[0])
	}

	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)

	flags.StringVar(&ctx.inputFilePath, "in", "", "input file path")
	flags.StringVar(&ctx.outputFilePath, "out", "", "output file path")

	flags.BoolVar(&ctx.noVerify, "no-verify", false, "disable verification")

	if err := flags.Parse(args[1:]); err != nil {
		return ctx, err
	}

	if ctx.inputFilePath == "" {
		return ctx, fmt.Errorf("input file path is required")
	}

	if ctx.outputFilePath == "" {
		return ctx, fmt.Errorf("output file path is required")
	}

	if flags.NArg() != 0 {
		return ctx, fmt.Errorf("unexpected arguments: %v", flags.Args())
	}

	return ctx, nil
}

func Run(args []string) error {
	ctx, err := readTermParams(args)
	if err != nil {
		return err
	}

	fmt.Printf("mode: %s\n", ctx.cryptoMode)
	fmt.Printf("input: %s\n", ctx.inputFilePath)
	fmt.Printf("output: %s\n", ctx.outputFilePath)
	fmt.Printf("no verify: %t\n", ctx.noVerify)

	return nil
}
