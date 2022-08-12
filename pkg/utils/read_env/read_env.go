package read_env

import (
	"fmt"
	"os"
)

// ReadOSEnv print the os env, set the option:"all" or set the env's key wanted.
func ReadOSEnv(option string) {
	if option == "all" || option == "" {
		environ := os.Environ()
		for i := range environ {
			fmt.Println(environ[i])
		}
	} else {
		envKey := os.Getenv(option)
		fmt.Printf("The OS ENV %s is %s\n", option, envKey)
	}
	return
}
