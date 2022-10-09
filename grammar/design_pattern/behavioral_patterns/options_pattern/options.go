package options_pattern

import "time"

// 选项模式有很多优点，例如：支持传递多个参数，并且在参数发生变化时保持兼容性；
//支持任意顺序传递参数；支持默认值；方便扩展；通过 WithXXX 的函数命名，可以使参数意义更加明确
// 结构体参数很多，创建结构体时，我们期望创建一个携带默认值的结构体变量，并选择性修改其中一些参数的值。
// 结构体参数经常变动，变动时我们又不想修改创建实例的函数。例如：结构体新增一个 retry 参数，
// 但是又不想在 NewConnect 入参列表中添加retry int这样的参数声明。
type Connect struct {
	addr  string
	cache bool
	time  time.Duration
}

const (
	defaultTime = 10
	defaltCach  = false
)

type options struct {
	timeout time.Duration
	caching bool
}

type Option interface {
	apply(*options)
}

type optionFunc func(*options)

func (f optionFunc) apply(o *options) {
	f(o)
}

func WithTimeOut(t time.Duration) Option {
	return optionFunc(func(o *options) {
		o.timeout = t
	})
}

func WithCaching(cache bool) Option {
	return optionFunc(func(o *options) {
		o.caching = cache
	})
}

func NewConnec(addr string, opts ...Option) (*Connect, error) {
	options := options{
		timeout: defaultTimeout,
		caching: defaultCaching,
	}

	for _, o := range opts {
		o.apply(&options)
	}

	return &Connect{
		addr:  addr,
		cache: options.caching,
		time:  options.timeout,
	}, nil
}
