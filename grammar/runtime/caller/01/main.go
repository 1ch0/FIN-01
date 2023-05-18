package main

import (
	"fmt"
	"path"
	"runtime"
)

func getCallerInfo(skip int) (info string) {

	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {

		info = "runtime.Caller() failed"
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	fileName := path.Base(file) // Base函数返回路径的最后一个元素
	return fmt.Sprintf("FuncName:%s, file:%s, line:%d ", funcName, fileName, lineNo)
}

func main() {

	// 打印出getCallerInfo函数自身的信息
	fmt.Println(getCallerInfo(0))
	// 打印出getCallerInfo函数的调用者的信息
	fmt.Println(getCallerInfo(1))
}
