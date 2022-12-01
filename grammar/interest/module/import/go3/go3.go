package go3

import (
	"fmt"
	"go-cx/grammar/interest/import/go2"
)

type ddd struct{}

func (d *ddd) Open() {
	fmt.Println("this is go3")
}

func init() {
	go2.Register("go3", &ddd{})
}
