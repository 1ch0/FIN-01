package read_env

import "testing"

func TestGetEnv(t *testing.T) {
	ReadOSEnv("GOVERSION")
	t.Log("--------------------------------------------------------------------")
	ReadOSEnv("")
	t.Log("--------------------------------------------------------------------")
	ReadOSEnv("all")
}
