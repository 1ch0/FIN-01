package main

import "fmt"

func every[E any](s []E, f func(E) bool) bool {
	for _, e := range s {
		if !f(e) {
			return false
		}
	}

	return true
}

func main() {
	allEven := every([]int{1, 2, 3, 4}, func(v int) bool {
		return v%2 == 0
	})
	if allEven {
		fmt.Println("All Even")
	}
}
