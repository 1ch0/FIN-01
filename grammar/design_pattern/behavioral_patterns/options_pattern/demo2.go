package options_pattern

import "time"

// 需要创建一个带默认值的选项，并用该选项创建实例
// 只需要实现一个函数来创建实例，但是也有缺点：为了创建 Connection 实例，每次我们都要创建 ConnectionOptions，操作起来比较麻烦
const (
	defaultTimeout = 10
	defaultCaching = false
)

type Connection struct {
	addr    string
	cache   bool
	timeout time.Duration
}

type ConnectionOptions struct {
	Caching bool
	Timeout time.Duration
}

func NewDefaultOptions() *ConnectionOptions {
	return &ConnectionOptions{
		Caching: defaultCaching,
		Timeout: defaultTimeout,
	}
}

func NewConnect(addr string, opts *ConnectionOptions) (*Connection, error) {
	return &Connection{
		addr:    addr,
		cache:   opts.Caching,
		timeout: opts.Timeout,
	}, nil
}
