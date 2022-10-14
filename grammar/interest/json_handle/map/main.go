package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	byt := []byte(`{
        "num":6.13,
        "strs":["a","b"],
        "obj":{"foo":{"bar":"zip","zap":6}}
    }`)
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"].(float64)
	fmt.Println(num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	obj := dat["obj"].(map[string]interface{})
	obj2 := obj["foo"].(map[string]interface{})
	fmt.Println(obj2)

}
