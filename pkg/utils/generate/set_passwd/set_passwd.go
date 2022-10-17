package set_passwd

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

func SetPasswd(length int) string {
	rand.Seed(time.Now().UnixNano())

	buf := make([]byte, length)
	buf[0] = Digits[rand.Intn(len(Digits))]
	buf[1] = Uppercase[rand.Intn(len(Uppercase))]
	buf[2] = Lowercase[rand.Intn(len(Lowercase))]
	buf[3] = Specials[rand.Intn(len(Specials))]

	for i := 4; i < length; i++ {
		buf[i] = All[rand.Intn(len(All))]
	}
	rand.Shuffle(len(buf), func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})
	passwd := string(buf) // E.g. "3i[g0|)z"
	return passwd
}
