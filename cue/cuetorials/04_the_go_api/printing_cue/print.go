package main

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
)

const val = `
i: int
s: string
t: [string]: string
_h: "hidden"
#d: foo: "bar"
`

func main() {
	var (
		c *cue.Context
		v cue.Value
	)

	c = cuecontext.New()

	v = c.CompileString(val)

	fmt.Printf("// %%v\n%v\n\n// %%# v\n%# v\n", v, v)
}
