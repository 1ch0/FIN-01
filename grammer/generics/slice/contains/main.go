package main

import "fmt"

func Contains[E comparable](s []E, v E) bool {
	for _, vs := range s {
		if v == vs {
			return true
		}
	}
	return false
}

func main() {
	vi := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(Contains(vi, 1))

	str := []string{"aa", "bb", "cc", "dd", "ee"}
	fmt.Println(Contains(str, "dd"))
}
