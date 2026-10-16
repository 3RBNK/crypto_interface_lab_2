package entropy

import (
	"crypto/rand"
	"crypto/sha256"
	"filecrypt/internal/app"
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
)

const (
	sampleCount = 256
	sampleDelay = 10 * time.Millisecond
	timeout     = 30 * time.Second
)

func generateSalt() ([]byte, error) {
	hash := sha256.New()

	random := make([]byte, app.SaltSize)
	if _, err := rand.Read(random); err != nil {
		return nil, err
	}

	hash.Write(random)

	fmt.Println("Move mouse to generate salt")

	globalX, globalY := robotgo.Location()
	deadline := time.Now().Add(timeout)

	count := 0

	for count < sampleCount {
		if time.Now().After(deadline) {
			return nil, fmt.Errorf("timeout of mouse move")
		}

		x, y := robotgo.Location()
		if x == globalX && y == globalY {
			time.Sleep(sampleDelay)
			continue
		}

		globalX, globalY = x, y

		count++

		fmt.Printf("\rCollected: %d/%d", count, sampleCount)

		time.Sleep(sampleDelay)
	}

	fmt.Println("\n Generate salt complete \n")
	return hash.Sum(nil), nil
}
