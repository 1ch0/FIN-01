package main

import "fmt"

func main() {
	var s []int
	if s == nil {
		fmt.Println("slice is empty")
	}
	// The nil slice is ready to use: no need to check for nil.
	s = append(s, 1)
	s = append(s, 2)
	fmt.Println(s)

	// A slice value with zero length is not the same as a nil slice
	// The length of nil is zero,
	// but a slice created using make([]T, 0) it has a length of zero but is not nil.
	var s1 []int
	var s2 = []int{}
	s3 := make([]int, 0)

	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)
	fmt.Println(s3 == nil)
	fmt.Println(len(s1))
	fmt.Println(len(s2))
	fmt.Println(len(s3))
}
