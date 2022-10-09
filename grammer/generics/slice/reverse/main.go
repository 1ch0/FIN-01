package main

import "fmt"

func Reverse[E any](s []E) []E {
	result := make([]E, 0, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		result = append(result, s[i])
	}
	return result
}

func main() {
	si := []int{1, 2, 3, 4, 5}
	fmt.Println(Reverse(si))

	str := []string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(Reverse(str))
}
