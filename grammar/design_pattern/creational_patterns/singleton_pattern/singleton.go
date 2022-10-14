package singleton_pattern

import "sync"

type singleton struct {
}

var ins *singleton
var once sync.Once

// 使用once.Do可以确保 ins 实例全局只被创建一次，
// once.Do 函数还可以确保当同时有多个创建动作时，只有一个创建动作在被执行，
// 内部也是 mutex.
func GetInsOr() *singleton {
	once.Do(func() {
		ins = &singleton{}
	})
	return ins
}
