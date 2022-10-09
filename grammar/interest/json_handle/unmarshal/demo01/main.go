package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	// 三方返回普通 json 字符串
	jsonRes := `{
  "Id": 1001,
  "Name": "frank"
 }`
	data := new(User)
	err := json.Unmarshal([]byte(jsonRes), &data)
	if err != nil {
		log.Printf("json Unmarshal err:%v\n", err)
		return
	}
	fmt.Printf("data=%+v", data)
}

type User struct {
	Id   int
	Name string
}
