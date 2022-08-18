package main

import (
	"fmt"

	"github.com/bytedance/sonic"
)

type User struct {
	Name  string `json:"user_name"`
	Age   int    `json:"age"`
	sex   string
	Work1 *Work1
	Work2 *Work2
	Work3 Work3
	Work4 interface{} `json:"work4"`
}

type Work1 struct {
	Name   string `json:"work1_name"`
	Salary float32
}

type Work2 struct {
	Name   string `json:"work2_name"`
	Salary float32
}

type Work3 struct {
	Name   string `json:"work3_name"`
	Salary float32
}

type Work4 struct {
	Name    string `json:"work4_name"`
	Salary  float32
	Address string `json:"work4_add"`
}

func main() {
	// 实例化User
	u1 := User{
		Name: "ares",
		Age:  18,
		sex:  "男",
	}
	// 指针
	w1 := Work1{
		Name:   "god1",
		Salary: 100,
	}
	u1.Work1 = &w1
	// 指针
	w2 := new(Work2)
	w2.Name = "god2"
	w2.Salary = 200
	u1.Work2 = w2
	// 非指针
	w3 := Work3{
		Name:   "god3",
		Salary: 300,
	}
	u1.Work3 = w3
	// 非指针
	w4 := Work4{
		Name:    "god4",
		Salary:  400,
		Address: "cbd",
	}
	u1.Work4 = w4

	jsonU, err := sonic.Marshal(u1)
	if err != nil {
		fmt.Println("生成json字符串错误")
	}

	// jsonU是[]byte类型，转化成string类型便于查看
	// {"user_name":"ares","age":18,"Work1":{"work1_name":"god1","Salary":100},"Work2":{"work2_name":"god2","Salary":200},"Work3":{"work3_name":"god3","Salary":300},"work4":{"work4_name":"god4","Salary":400,"work4_add":"cbd"}}
	fmt.Println(string(jsonU))
}
