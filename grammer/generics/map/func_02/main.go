package main

import (
	"fmt"
	"strings"
)

type mapFunc[E any] func(E) E

func Map[E any](s []E, f mapFunc[E]) []E {
	result := make([]E, len(s))
	for i := range s {
		result[i] = f(s[i])
	}
	return result
}

func main() {
	s := []string{"a", "b", "c", "d", "e"}
	fmt.Println(Map(s, strings.ToUpper))
}
