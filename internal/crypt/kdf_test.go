package crypt

import (
	"bytes"
	"testing"
)

func TestDeriveKeys(t *testing.T) {
	password := "my-secret-password"
	salt := bytes.Repeat([]byte{0x42}, 32)

	encKey1, macKey1, err := deriveKey(password, salt)
	if err != nil {
		t.Fatal(err)
	}

	encKey2, macKey2, err := deriveKey(password, salt)
	if err != nil {
		t.Fatal(err)
	}

	if len(encKey1) != 32 || len(macKey1) != 32 {
		t.Fatal("invalid key size")
	}

	if !bytes.Equal(encKey1, encKey2) ||
		!bytes.Equal(macKey1, macKey2) {
		t.Fatal("keys are not deterministic")
	}

	if bytes.Equal(encKey1, macKey1) {
		t.Fatal("encryption and MAC keys must differ")
	}
}
