package main

import "github.com/tidwall/sjson"

const JSON = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`

func main() {
	value, _ := sjson.Set(JSON, "name.last", "Anderson")
	println(value)
}
