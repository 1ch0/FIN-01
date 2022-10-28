package _1_facade

import "fmt"

type Test interface {
	Test() string
}

func NewAPI() Test {
	return &apiImpl{
		a: NewTestA(),
		b: NewTestB(),
	}
}

type apiImpl struct {
	a AModuleImpl
	b BModuleImpl
}

func (a *apiImpl) Test() string {
	aRet := a.a.Test()
	bRet := a.b.Test()
	return fmt.Sprintf("%s\n%s\n", aRet, bRet)
}

type AModuleImpl interface {
	Test() string
}

func NewTestA() AModuleImpl {
	return &testAImpl{}
}

type testAImpl struct {
}

func (a *testAImpl) Test() string {
	return "testA"
}

type BModuleImpl interface {
	Test() string
}

func NewTestB() AModuleImpl {
	return &testBImpl{}
}

type testBImpl struct {
}

func (b *testBImpl) Test() string {
	return "testB"
}
