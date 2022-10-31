package main

import (
	"fmt"
	"time"
)

func main() {
	gen := func() <-chan int {
		dst := make(chan int)
		go func() {
			var n int
			for { // goroutine 泄露
				dst <- n
				n++
			}
		}()
		return dst
	}
	for n := range gen() {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
	time.Sleep(1 * time.Second)
}
