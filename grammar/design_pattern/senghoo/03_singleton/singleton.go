package _3_singleton

import "sync"

type Singleton interface {
	foo()
}

type singleton struct{}

func (receiver singleton) foo() {

}

var (
	instance *singleton
	once     sync.Once
)

func GetInstance() Singleton {
	once.Do(func() {
		instance = &singleton{}
	})

	return instance
}
