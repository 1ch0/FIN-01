package main

import "fmt"

func main() {
	unsuccessNum := make([]int, 0)
	for i := 0; i < 5; i++ {
		unsuccessNum = append(unsuccessNum, i)
		i += 1
	}
	fmt.Println(unsuccessNum)
	for _, v := range unsuccessNum {
		fmt.Println(v)
	}
}
