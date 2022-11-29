package main

import (
	"go-cx/grammar/interest/module/import/go2"
	_ "go-cx/grammar/interest/module/import/go3"
)

func main() {
	go2.DDD("go3")
}
