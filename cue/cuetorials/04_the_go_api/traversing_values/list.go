package main

import (
	"fmt"
	"os"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
)

func main() {
	c := cuecontext.New()
	d, _ := os.ReadFile("value.cue")
	val := c.CompileBytes(d)

	val = val.LookupPath(cue.ParsePath("obj.list"))

	iter, _ := val.List()

	for iter.Next() {
		fmt.Println(iter.Value())
	}
}
