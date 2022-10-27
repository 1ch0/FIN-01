package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	defer func() {
		fmt.Println("goroutines: ", runtime.NumGoroutine())
	}()
	GetData()
	time.Sleep(time.Second * 2)
}

func GetData() {
	var ch chan struct{}
	go func() {
		<-ch
	}()
}
