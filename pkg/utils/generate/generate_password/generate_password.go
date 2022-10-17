package generate_password

import (
	"math/rand"
	"time"
)

const (
	Digits    = "0123456789"
	Uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Lowercase = "abcdefghijklmnopqrstuvwxyz"
	Specials  = "~=+%^*/()[]{}/!@#$?|"
	All       = Digits + Uppercase + Lowercase + Specials
)

// GeneratePassword generate password, length to set the password length
// level: default(contains number&character), number(only contains number), char(only contains character),
// mix(contains number&character), advance(contains number&character&specials)
func GeneratePassword(length int, level string) string {
	rand.Seed(time.Now().UnixNano())

	buf := make([]byte, length)
	levelNum := 0
	switch level {
	case "num":
		buf[0] = Digits[rand.Intn(len(Digits))]
		levelNum = 1
	case "char":
		buf[0] = Uppercase[rand.Intn(len(Uppercase))]
		buf[1] = Lowercase[rand.Intn(len(Lowercase))]
		levelNum = 2
	case "mix":
		buf[0] = Digits[rand.Intn(len(Digits))]
		buf[1] = Uppercase[rand.Intn(len(Uppercase))]
		buf[2] = Lowercase[rand.Intn(len(Lowercase))]
		levelNum = 3
	case "advance":
		buf[0] = Digits[rand.Intn(len(Digits))]
		buf[1] = Uppercase[rand.Intn(len(Uppercase))]
		buf[2] = Lowercase[rand.Intn(len(Lowercase))]
		buf[3] = Specials[rand.Intn(len(Specials))]
		levelNum = 4
	default:
		buf[0] = Digits[rand.Intn(len(Digits))]
		buf[1] = Uppercase[rand.Intn(len(Uppercase))]
		buf[2] = Lowercase[rand.Intn(len(Lowercase))]
		levelNum = 3
	}

	for i := levelNum; i < length; i++ {
		buf[i] = All[rand.Intn(len(All))]
	}
	rand.Shuffle(len(buf), func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})
	passwd := string(buf) // E.g. "3i[g0|)z"
	return passwd
}
