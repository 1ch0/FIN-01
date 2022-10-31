package singleton_pattern

import "sync"

// 懒汉方式是开源项目中使用最多的，
// 但它的缺点是非并发安全，在实际使用时需要加锁
// 在创建 ins 时，如果 ins==nil，就会再创建一个 ins 实例，这时候单例就会有多个实例
// 为了解决懒汉方式非并发安全的问题，需要对实例进行加锁
type singleton3 struct {
}

var ins3 *singleton3
var mu sync.Mutex

func GetIns() *singleton3 {
	if ins == nil {
		mu.Lock()
		if ins == nil {
			ins3 = &singleton3{}
		}
		mu.Unlock()
	}
	return ins3
}
