package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
 * @Description
 * @Author guirongguo
 * @Email 3095764372@qq.com
 * @Date 2021/8/30 23:00
 **/

// 简单使用
func TestSimpleRand(t *testing.T) {
	t.Log("start ...")
	assert := assert.New(t)
	assert.Equal(1, 1)
	assert.NotEqual(1, 2)
	assert.NotNil("123")
	assert.IsType([]string{}, []string{""})

	assert.Contains("Hello World", "World")
	assert.Contains(map[string]string{"Hello": "World"}, "Hello")
	assert.Contains([]string{"Hello", "World"}, "Hello")
	assert.True(true)
	assert.True(false)
	t.Log("next ...")
	var s []string
	assert.Empty(s)
	assert.Nil(s)
	t.Log("end ...")
}

// 一般用的更多的是表驱动方式把同一个单元的测试用例都放在一起
func TestCalculate(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-1, 1},
		{0, 2},
		{-5, -3},
		{99999, 100001},
	}

	for _, test := range tests {
		assert.Equal(Calculate(test.input), test.expected)
	}
}
