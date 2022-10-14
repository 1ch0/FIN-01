package main

import (
	"encoding/json"
	"fmt"
)

// UserInfo 用户信息
type UserInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	u1 := UserInfo{Name: "custom", Age: 18}

	b, _ := json.Marshal(&u1)
	var m map[string]interface{}
	_ = json.Unmarshal(b, &m)
	for k, v := range m {
		fmt.Printf("key:%v value:%v\n\n", k, v)
		fmt.Printf("key:%v value:%v value type:%T\n", k, v, v)
	}
}
