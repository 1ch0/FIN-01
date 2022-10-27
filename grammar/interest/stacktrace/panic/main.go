package main

import "time"

// 如果想让它把所有的goroutine信息都输出出来，可以设置 GOTRACEBACK=1:
// GOTRACEBACK=1 go run p.go
func main() {
	go a()
	m1()
}

func m1() {
	m2()
}

func m2() {
	m3()
}

func m3() {
	panic("panic from m3")
}

func a() {
	time.Sleep(time.Hour)
}
