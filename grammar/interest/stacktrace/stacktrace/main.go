package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

func main() {
	setupSigusr1Trap()
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
func setupSigusr1Trap() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGUSR1)
	go func() {
		for range c {
			DumpStacks()
		}
	}()
}
func DumpStacks() {
	buf := make([]byte, 16384)
	buf = buf[:runtime.Stack(buf, true)]
	fmt.Printf("=== BEGIN goroutine stack dump ===\n%s\n=== END goroutine stack dump ===", buf)
}
