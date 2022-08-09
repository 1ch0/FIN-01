package main

import (
	"fmt"

	"cuelang.org/go/cue"

	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/errors"
)

const val = `
i: int
s: string
t: [string]: string
_h: int
_h: "hidden"
#d: int
#d: "bar"
`

func main() {
	c := cuecontext.New()
	v := c.CompileString(val)
	if err := v.Validate(
		cue.Concrete(false),
		cue.Definitions(false),
		cue.Hidden(false),
		cue.Optional(false),
	); err != nil {
		printErr("loose error", err)
	}

	printErr("loose error", loose(v))
	printErr("every error", every(v))
	printErr("strict error", strict(v))

	fmt.Printf("\nvalue:\n%#v\n", v)
}

func printErr(prefix string, err error) {
	if err != nil {
		msg := errors.Details(err, nil)
		fmt.Printf("%s:\n%s\n", prefix, msg)
	}
}

func loose(v cue.Value) error {
	return v.Validate(
		cue.Concrete(false),
		cue.Definitions(false),
		cue.Hidden(false),
		cue.Optional(false),
	)
}

func every(v cue.Value) error {
	return v.Validate(
		cue.Concrete(false),
		cue.Definitions(true),
		cue.Hidden(true),
		cue.Optional(true),
	)
}

func strict(v cue.Value) error {
	return v.Validate(
		cue.Final(),
		cue.Concrete(false),
		cue.Definitions(true),
		cue.Hidden(true),
		cue.Optional(true),
	)
}
