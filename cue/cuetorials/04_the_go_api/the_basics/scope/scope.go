package main

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
)

const schema = `
#schema: {
	i: int
	s: string
}
`

const val = `
v: #schema & {
	i: 1
	s: "hello"
}
`

func main() {
	var (
		c *cue.Context
		s cue.Value
		v cue.Value
	)

	c = cuecontext.New()

	s = c.CompileString(schema)

	v = c.CompileString(val, cue.Scope(s))

	fmt.Println(v)
}
