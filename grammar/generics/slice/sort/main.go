package main

import (
	"fmt"
	"sort"

	"golang.org/x/exp/constraints"
)

func Sort[E constraints.Ordered](s []E) []E {
	result := make([]E, len(s))
	copy(result, s)
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	return result
}

func main() {
	fmt.Println(Sort([]string{"a", "c", "b"}))

	fmt.Println(Sort([]int{1, 22, 33, 4}))
}
