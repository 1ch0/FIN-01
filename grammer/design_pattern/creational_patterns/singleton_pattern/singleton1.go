package singleton_pattern

// 饿汉方式的单例模式代码
// 实例是在包被导入时初始化的，所以如果初始化耗时，会导致程序加载时间比较长
type singleton1 struct {
}

var ins1 *singleton1 = &singleton1{}

func GetIns1Or() *singleton1 {
	return ins1
}
