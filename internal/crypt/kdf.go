package crypt

import (
	"crypto/pbkdf2"
	"crypto/sha256"
	"filecrypt/internal/app"
	"fmt"
)

func deriveKey(password string, salt []byte) (encryptionKey, macKey []byte, err error) {
	if len(salt) != app.SaltSize {
		return nil, nil, fmt.Errorf("salt size is invalid: %d", len(salt))
	}

	if password == "" {
		return nil, nil, fmt.Errorf("password cannot be empty")
	}

	derived, err := pbkdf2.Key(
		sha256.New,
		password,
		salt,
		app.Iterations,
		2*app.KeySize,
	)

	if err != nil {
		return nil, nil, fmt.Errorf("deriving key: %w", err)
	}

	encryptionKey = derived[:app.KeySize]
	macKey = derived[app.KeySize:]

	return encryptionKey, macKey, nil
}
