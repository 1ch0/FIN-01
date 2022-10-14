package main

import "github.com/tidwall/gjson"

const JSON = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`

var jsonRes = `{
        "num":6.13,
        "strs":["a","b"],
        "obj":{"foo":{"bar":"zip","zap":6}}
    }`

func main() {
	value := gjson.Get(JSON, "name.last")
	println(value.String())

	value = gjson.Get(jsonRes, "obj.foo.bar")
	println(value.String())
}
