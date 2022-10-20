package main

import "fmt"

type Bar struct {
	n    int
	f    float64
	next *Bar
}

func main() {
	fmt.Println([2]Bar{})
}
