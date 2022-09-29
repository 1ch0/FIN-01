package main

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
)

type Val struct {
	I int    `json:"i"`
	S string `json:"s,omitempty"`
	b bool
}

func main() {
	var (
		c *cue.Context
		v cue.Value
	)

	val := Val{
		I: 1,
		S: "hello",
		b: true,
	}
	c = cuecontext.New()

	v = c.Encode(val)

	fmt.Println(v)

	t := c.EncodeType(Val{})
	fmt.Println(t)
}
