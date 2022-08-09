package main

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
)

const val = `
i: int
s: "hello"
`

func main() {
	var (
		c *cue.Context
		v cue.Value
	)

	c = cuecontext.New()

	v = c.CompileString(val)

	fmt.Println(v)
}
