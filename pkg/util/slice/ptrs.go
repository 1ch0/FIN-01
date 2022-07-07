package main

import "fmt"

const MAX = 3

var numbers []int = []int{10, 100, 2 << 2}

func correct() {
	// 指针数组
	var ptrs [MAX]*int
	// 将numbers数组中的每一个元素的地址
	// 赋值给指针数组中的每一个元素
	for i := 0; i < MAX; i++ {
		ptrs[i] = &numbers[i]
	}

	for key, value := range ptrs {
		// 每个元素的内存地址是不同的
		fmt.Printf("index: %d value: %d, address of value: %d\n", key, *value, value)
	}
}

func mistake() {
	var ptrs [MAX]*int
	for k, v := range numbers {
		ptrs[k] = &v
	}

	for k, v := range ptrs {
		fmt.Printf("index: %d value: %d, address of value: %d\n", k, *v, v)
	}
}

func main() {
	correct()
	mistake()
}
