package cli

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

func ReadPassword() (string, error) {
	fd := int(os.Stdin.Fd())

	if !term.IsTerminal(fd) {
		return "", fmt.Errorf("stdin is not a terminal")
	}

	fmt.Printf("Enter password: ")

	password, err := term.ReadPassword(fd)
	if err != nil {
		return "", fmt.Errorf("failed to read password: %w", err)
	}

	if len(password) == 0 {
		return "", fmt.Errorf("password is empty")
	}

	return string(password), nil
}
