package format

import "fmt"

const (
	SaltSize   = 32
	HMACSize   = 32
	IVSize     = 16
	KeySize    = 32
	Iterations = 600_000
	HeaderSize = SaltSize + HMACSize + IVSize
)

func Pack(salt, mac, iv, ciphertext []byte) ([]byte, error) {
	if len(salt) != SaltSize {
		return nil, fmt.Errorf("invalid salt size: %d", len(salt))
	}

	if len(mac) != HMACSize {
		return nil, fmt.Errorf("invalid HMAC size: %d", len(mac))
	}

	if len(iv) != IVSize {
		return nil, fmt.Errorf("invalid IV size: %d", len(iv))
	}

	data := make([]byte, 0, HeaderSize+len(ciphertext))

	data = append(data, salt...)
	data = append(data, mac...)
	data = append(data, iv...)
	data = append(data, ciphertext...)

	return data, nil
}

func Unpack(data []byte) (salt, mac, iv, ciphertext []byte, err error) {
	if len(data) < HeaderSize {
		err = fmt.Errorf("encrypted file is too short")
		return
	}

	salt = data[:SaltSize]

	mac = data[SaltSize : SaltSize+HMACSize]

	iv = data[SaltSize+HMACSize : HeaderSize]

	ciphertext = data[HeaderSize:]

	return
}
