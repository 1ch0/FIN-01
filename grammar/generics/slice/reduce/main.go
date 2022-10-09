package main

import "fmt"

type reduceFunc[E any] func(E, E) E

func Reduce[E any](s []E, init E, f reduceFunc[E]) E {
	cur := init
	for _, v := range s {
		cur = f(cur, v)
	}
	return cur
}

func main() {
	s := []int{1, 2, 3, 4}
	sum := Reduce(s, 0, func(cur, next int) int {
		return cur + next
	})
	fmt.Println(sum)

	p := Reduce(s, 1, func(cur, next int) int {
		return cur * next
	})
	fmt.Println(p)

	s1 := []string{"a", "b", "c", "d", "e"}
	j := Reduce(s1, "", func(c, n string) string {
		return c + ":" + n
	})
	fmt.Println(j)
}
