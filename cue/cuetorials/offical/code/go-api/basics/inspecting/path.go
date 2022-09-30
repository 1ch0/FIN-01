package main

import (
	"fmt"
	"os"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
)

func main() {
	c := cuecontext.New()

	// read and compile value
	d, _ := os.ReadFile("value.cue")
	val := c.CompileBytes(d)

	paths := []string{
		"a",
		"d.f",
		"l",
	}

	for _, path := range paths {
		fmt.Printf("====  %s  ====\n", path)
		v := val.LookupPath(cue.ParsePath(path))
		p := v.Path()
		fmt.Printf("%q\n%# v\n", p, v)
	}
}
