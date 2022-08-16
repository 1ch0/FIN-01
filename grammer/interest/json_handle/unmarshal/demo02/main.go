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
	data2 := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonRes), &data2)
	if err != nil {
		log.Printf("json Unmarshal err:%v\n", err)
		return
	}
	fmt.Printf("data2=%+v", data2)
}
