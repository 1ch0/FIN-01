
一个能检测 Go 协程泄露隐患的工具库
https://mp.weixin.qq.com/s/XrhdU95CswLGwS0CLxqZmg

使用runtime.Stack()方法获取当前运行的所有goroutine的栈信息，
默认定义不需要检测的过滤项，默认定义检测次数+检测间隔，不断周期进行检测，
最终在多次检查后仍没有找到剩下的goroutine则判断没有发生goroutine泄漏。