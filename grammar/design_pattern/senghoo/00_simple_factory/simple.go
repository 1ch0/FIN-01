package _0_simple_factory

import "fmt"

type API interface {
	Say(name string) string
}

func NewAPI(name string) API {
	switch name {
	case "hi":
		return &hiAPI{}
	case "hello":
		return &helloAPI{}
	default:
		return &hiAPI{}
	}
}

type hiAPI struct {
}

func (receiver hiAPI) Say(name string) string {
	return fmt.Sprintf("Hi %s", name)
}

type helloAPI struct {
}

func (receiver helloAPI) Say(name string) string {
	return fmt.Sprintf("Hello %s", name)
}
