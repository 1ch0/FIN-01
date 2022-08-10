package main

import (
	"go-cx/grammer/interest/for/a"
	"go-cx/grammer/interest/for/b"
)

var name = "for 函数式编程"

func main() {
	Setup(name)
}

func Setup(name string) error {
	for _, setup := range []func(name string) error{
		a.Setup, b.Setup,
	} {
		if err := setup(name); err != nil {
			return err
		}
	}
	return nil
}
