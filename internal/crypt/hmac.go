package crypt

import (
	"crypto/hmac"
	"crypto/sha256"
)

func GenerateHMAC(data, key []byte) []byte {
	hash := hmac.New(sha256.New, key)
	hash.Write(data)

	return hash.Sum(nil)
}

func VerifyHMAC(data, key, excepted []byte) bool {
	actualMAC := GenerateHMAC(data, key)
	return hmac.Equal(actualMAC, excepted)
}
