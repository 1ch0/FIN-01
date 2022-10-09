package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	// 三方返回嵌套 json 字符串
	jsonRes := `{
 "Id": 1001,
 "Name": "frank",
 "Details": {
 "Gender": "man",
 "Age": 18,
 "Phone": "13800138000",
 "address": "Beijing"
 }
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
	Id      int
	Name    string
	Details Details
}

type Details struct {
	Gender  string
	Age     int
	Phone   string
	Address string
}
