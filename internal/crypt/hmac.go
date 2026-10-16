package crypt

import (
	"crypto/hmac"
	"crypto/sha256"
)

func generateHMAC(data, key []byte) []byte {
	hmac := hmac.New(sha256.New, key)
	hmac.Write(data)

	return hmac.Sum(nil)
}

func verifyHMAC(data, key, excepted []byte) bool {
	actualMAC := generateHMAC(data, key)
	return hmac.Equal(actualMAC, excepted)
}
