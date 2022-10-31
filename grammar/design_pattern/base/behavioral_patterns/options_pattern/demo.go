package options_pattern

import "time"

// 分别开发两个用来创建实例的函数，一个可以创建带默认值的实例，一个可以定制化创建实例
// 使用这种方式，创建同一个 Connection 实例，却要实现两个不同的函数，实现方式很不优雅。

const (
	defaultTimeout1 = 10
	defaultCaching1 = false
)

type Connection1 struct {
	addr    string
	cache   bool
	timeout time.Duration
}

func NewConnect1(addr string) (*Connection1, error) {
	return &Connection1{
		addr:    addr,
		cache:   defaultCaching1,
		timeout: defaultTimeout1,
	}, nil
}

func NewConnectWithOptions(addr string, cache bool, time time.Duration) (*Connection1, error) {
	return &Connection1{
		addr:    addr,
		cache:   cache,
		timeout: time,
	}, nil
}
