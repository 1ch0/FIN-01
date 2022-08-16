package main

import (
	"encoding/json"
	"fmt"
)

type ourData struct {
	Num  float64                      `json:"num"`
	Strs []string                     `json:"strs"`
	Obj  map[string]map[string]string `json:"obj"`
}

func main() {
	byt := []byte(`{
        "num":6.13,
        "strs":["a","b"],
        "obj":{"foo":{"bar":"zip","zap":6}}
    }`)

	res := ourData{}
	json.Unmarshal(byt, &res)
	fmt.Println(res.Num)
	fmt.Println(res.Strs)
	fmt.Println(res.Obj)
}
