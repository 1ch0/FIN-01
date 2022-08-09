package main

import "C"
import "fmt"

func main() {
	a := 0
	next := map[string]interface{}{
		"bar": a + 1,
	}
	for i := 0; i < 10; i++ {
		next["hello"] = "world"
	}

	fmt.Println(next)
}
