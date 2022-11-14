package main

import (
	"fmt"
)

//一个 rune 字面量代表一个 rune 常量。而常量分为有类型常量（typed）和无类型常量（untyped）。
//而字面量属于无类型常量，
//只不过每一个无类型常量都有一个默认类型。比如 a 字面量是一个无类型常量，它的默认类型是 rune。
//当在上下文中需要一个无类型常量带类型的值时，会进行隐式转换（或使用默认类型）
// a constant index that is untyped is given type int

//func main() {
//	m := [...]int{
//		97: 1,
//		98: 2,
//		99: 3,
//	}
//	m[97] = 3
//	fmt.Println(len(m))
//}

func main() {
	m := [...]int{
		'a': 1,
		'b': 2,
		'c': 3,
	}
	m['a'] = 3
	fmt.Println(len(m))
}
