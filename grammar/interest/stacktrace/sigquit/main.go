package main

import "time"

// 运行kill -SIGQUIT <pid> 杀死这个程序，程序在退出的时候输出strack trace
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
	time.Sleep(time.Hour)
}

func a() {
	time.Sleep(time.Hour)
}
