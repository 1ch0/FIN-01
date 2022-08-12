package main

import (
	"fmt"
	"os"
)

func main() {
	environ := os.Environ()
	for i := range environ {
		fmt.Println(environ[i])
	}
	fmt.Println("**************************")
	goPath := os.Getenv("VERSION")
	fmt.Printf("GOPATH is %s\n", goPath)
}
