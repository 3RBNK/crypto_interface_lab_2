package crypt

import (
	"testing"
)

func TestHMAC(t *testing.T) {
	key := []byte("secret-key")
	data := []byte("Hello, World!")

	mac := generateHMAC(data, key)

	if len(mac) != sha256Size {
		t.Fatal("invalid HMAC size")
	}

	if !verifyHMAC(data, key, mac) {
		t.Fatal("HMAC verification failed")
	}

	modified := []byte("Hello, World?")

	if verifyHMAC(modified, key, mac) {
		t.Fatal("modified data passed verification")
	}
}

const sha256Size = 32
