package generate

import (
	"math/rand"
	"time"
)

const (
	defaultTokenLen int = 16
)

func GenerateToken() string {
	rand.Seed(time.Now().UnixNano())
	runes := []rune("abcdefghijklmnopqrstuvwxyz0123456789")

	b := make([]rune, defaultTokenLen)
	for i := range b {
		b[i] = runes[rand.Intn(len(runes))] // #nosec
	}
	return string(b)
}
