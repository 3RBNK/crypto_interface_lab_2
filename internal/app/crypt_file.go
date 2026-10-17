package app

import (
	"crypto/rand"
	"filecrypt/internal/crypt"
	"filecrypt/internal/entropy"
	"filecrypt/internal/format"
	"fmt"
)

func EncryptFile(inputPath, outputPath, password string) error {
	plainText, err := readBinFile(inputPath)
	if err != nil {
		return fmt.Errorf("failed to read input file: %w", err)
	}

	salt, err := entropy.GenerateSalt()
	if err != nil {
		return fmt.Errorf("failed to generate salt: %w", err)
	}

	encKey, macKey, err := crypt.DeriveKey(password, salt)
	if err != nil {
		return fmt.Errorf("failed to derive keys: %w", err)
	}

	iv := make([]byte, format.IVSize)
	if _, err := rand.Read(iv); err != nil {
		return fmt.Errorf("failed to generate IV: %w", err)
	}

	cipherText, err := crypt.EncryptCFB(plainText, encKey, iv)
	if err != nil {
		return fmt.Errorf("failed to encrypt file: %w", err)
	}

	mac := crypt.GenerateHMAC(plainText, macKey)

	packedData, err := format.Pack(salt, mac, iv, cipherText)
	if err != nil {
		return fmt.Errorf("failed to pack data: %w", err)
	}

	if err := writeBinFile(outputPath, packedData); err != nil {
		return fmt.Errorf("failed to write output file: %w", err)
	}

	return nil
}

func DecryptFile(inputPath, outputPath string, password string, verifyMac bool) error {
	cipherFile, err := readBinFile(inputPath)
	if err != nil {
		return fmt.Errorf("failed to read input file: %w", err)
	}

	salt, mac, iv, cipherText, err := format.Unpack(cipherFile)
	if err != nil {
		return fmt.Errorf("failed to unpack input file: %w", err)
	}

	encKey, macKey, err := crypt.DeriveKey(password, salt)
	if err != nil {
		return fmt.Errorf("failed to derive keys: %w", err)
	}

	plainText, err := crypt.DecryptCFB(cipherText, encKey, iv)
	if err != nil {
		return fmt.Errorf("failed to decrypt file: %w", err)
	}

	if verifyMac && !crypt.VerifyHMAC(plainText, macKey, mac) {
		return fmt.Errorf("failed to verify mac")
	}

	if err = writeBinFile(outputPath, plainText); err != nil {
		return fmt.Errorf("failed to write output file: %w", err)
	}

	return nil
}
