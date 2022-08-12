package generate_password

import (
	"testing"

	"gotest.tools/assert"
)

func TestGeneratePassword(t *testing.T) {
	t.Log(GeneratePassword(10, ""))
	assert.Equal(t, len(GeneratePassword(10, "")), 10)

	t.Log(GeneratePassword(15, "num"))
	t.Log(GeneratePassword(20, "char"))
	t.Log(GeneratePassword(20, "mix"))
	t.Log(GeneratePassword(30, "advance"))

}
