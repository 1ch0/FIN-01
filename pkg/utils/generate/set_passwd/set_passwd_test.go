package set_passwd

import (
	"math/rand"
	"testing"
	"time"
)

func TestGetPasswd(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	length := rand.Intn(20)
	passwd := SetPasswd(length)
	t.Logf("passwd: %s", passwd)
}
