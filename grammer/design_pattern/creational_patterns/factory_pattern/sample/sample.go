package sample

import "fmt"

// 简单工厂模式是最常用、最简单的。
//它就是一个接受一些参数，然后返回 Person 实例的函数
type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() {
	fmt.Printf("Hi! I'm %s", p.Name)
}

func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}
