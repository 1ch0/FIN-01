package options_pattern

import (
	"testing"
)

func TestNewConnec(t *testing.T) {
	c, _ := NewConnec("a", WithTimeOut(15), WithCaching(true))
	t.Log(c)
}
