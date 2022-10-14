package main

import "fmt"

type filterFunc[E any] func(E) bool

func Filter[E any](a []E, f filterFunc[E]) []E {
	var n []E
	for _, e := range a {
		if f(e) {
			n = append(n, e)
		}
	}
	return n
}

func main() {
	vi := []int{1, 2, 3, 4, 5, 6}
	vi = Filter(vi, func(v int) bool {
		return v < 4
	})
	fmt.Println(vi)
}
