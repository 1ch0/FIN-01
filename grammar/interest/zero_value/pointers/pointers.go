package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) IsAdult() bool {
	if p == nil {
		return false
	}
	return p.Age >= 18
}

func main() {
	p := &Person{
		"aaa",
		11,
	}
	if p.IsAdult() {
		fmt.Println("Adult")
	} else {
		fmt.Println("Minor")
	}
}
