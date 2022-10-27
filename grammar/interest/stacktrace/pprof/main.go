package main

import (
	"os"
	"runtime/pprof"
	"time"
)

// 通过debug.PrintStack()方法可以将当前所在的goroutine的stack trace打印出来
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
	pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)
	time.Sleep(time.Hour)
}

func a() {
	time.Sleep(time.Hour)
}
