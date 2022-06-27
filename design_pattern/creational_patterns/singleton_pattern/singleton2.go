package singleton_pattern

// 懒汉方式是开源项目中使用最多的，
// 但它的缺点是非并发安全，在实际使用时需要加锁
// 在创建 ins 时，如果 ins==nil，就会再创建一个 ins 实例，这时候单例就会有多个实例
// 为了解决懒汉方式非并发安全的问题，需要对实例进行加锁
type singleton2 struct {
}

var ins2 *singleton2

func GetIns2Or() *singleton2 {
	if ins2 == nil {
		ins2 = &singleton2{}
	}

	return ins2
}
