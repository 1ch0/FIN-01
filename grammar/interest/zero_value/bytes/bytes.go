package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var b bytes.Buffer

	b.Write([]byte("Hello"))
	b.Write([]byte(" "))
	b.Write([]byte("World"))

	fmt.Println(b.String())

	io.Copy(os.Stdout, &b)
}
