package app

import (
	"io"
	"os"
)

func readBinFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	return io.ReadAll(f)
}

func writeBinFile(path string, data []byte) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.Write(data)

	return err
}
