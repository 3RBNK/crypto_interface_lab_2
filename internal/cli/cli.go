package cli

import (
	"filecrypt/internal/app"
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
		PrintUsage()
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

	if ctx.inputFilePath == ctx.outputFilePath {
		return fmt.Errorf("input and output file cannot be the same file")
	}

	if ctx.cryptoMode == Encrypt && ctx.noVerify {
		return fmt.Errorf("--no-verify cannot be used with encrypt")
	}

	password, err := ReadPassword()
	if err != nil {
		return err
	}

	switch ctx.cryptoMode {
	case Encrypt:
		err = app.EncryptFile(ctx.inputFilePath, ctx.outputFilePath, password)
	case Decrypt:
		if ctx.noVerify {
			fmt.Println("Warning: HMAC verification is disabled")
		}

		err = app.DecryptFile(ctx.inputFilePath, ctx.outputFilePath, password, !ctx.noVerify)
	default:
		PrintUsage()
		return fmt.Errorf("unknown crypto mode: %d", ctx.cryptoMode)
	}

	if err != nil {
		return err
	}

	fmt.Println("Operation completed")
	return nil
}

func PrintUsage() {
	fmt.Println(`Usage:
  filecrypt <mode> --in <input> --out <output> [options]

Modes:
  encrypt        Encrypt input file
  decrypt        Decrypt input file

Arguments:
  --in           Path to input file
  --out          Path to output file

Options:
  --no-verify    Disable HMAC verification (decrypt only)
  -h, --help     Show this help message

Examples:
  filecrypt encrypt --in plain.txt --out encrypted.bin
  filecrypt decrypt --in encrypted.bin --out plain.txt
  filecrypt decrypt --in encrypted.bin --out plain.txt --no-verify

The password is entered securely after starting the program.`)
}
