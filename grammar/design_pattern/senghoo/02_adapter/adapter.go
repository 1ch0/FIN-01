package _2_adapter

import "fmt"

type Target interface {
	Request() string
}

func NewAdapter() Target {
	return adapterImpl{
		NewAdaptee(),
	}
}

type adapterImpl struct {
	adapteeImpl
}

func (receiver adapterImpl) Request() string {
	return receiver.SpecialRequest()
}

type Adaptee interface {
	SpecialRequest() string
}

func NewAdaptee() adapteeImpl {
	return adapteeImpl{}
}

type adapteeImpl struct {
}

func (receiver adapteeImpl) SpecialRequest() string {
	return fmt.Sprintf("adaptee")
}
