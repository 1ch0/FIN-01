package factory

import (
	"testing"
)

func TestNewPersonFactory(t *testing.T) {
	newBaby := NewPersonFactory(1)
	baby := newBaby("john")
	t.Log(baby)

	newTeenager := NewPersonFactory(16)
	teen := newTeenager("jill")
	t.Log(teen)
}
