package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"filecrypt/internal/app"
	"fmt"
)

func encryptCFB(plaintext, key, iv []byte) ([]byte, error) {
	if len(key) != app.KeySize {
		return nil, fmt.Errorf("invalid AES-256 key size: expected 32 bytes, got %d", len(key))
	}
	if len(iv) != aes.BlockSize {
		return nil, fmt.Errorf("invalid IV size: expected %d bytes, got %d", aes.BlockSize, len(iv))
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("create AES cipher: %w", err)
	}

	ciphertext := make([]byte, len(plaintext))
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext, plaintext)

	return ciphertext, nil
}

func decryptCFB(ciphertext, key, iv []byte) ([]byte, error) {
	if len(key) != 32 {
		return nil, fmt.Errorf("invalid AES-256 key size: expected 32 bytes, got %d", len(key))
	}
	if len(iv) != aes.BlockSize {
		return nil, fmt.Errorf("invalid IV size: expected %d bytes, got %d", aes.BlockSize, len(iv))
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("create AES cipher: %w", err)
	}

	plaintext := make([]byte, len(ciphertext))
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(plaintext, ciphertext)

	return plaintext, nil
}
