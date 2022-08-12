package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// json字符串解析时，需要一个“接收体”接受解析后的数据，且Unmarshal时接收体必须传递指针。
//解析时，接收体可自行定义。json串中的key自动在接收体中寻找匹配的项进行赋值。
//匹配规则：先查找与key一样的json标签，找到则赋值给该标签对应的变量；没有json标签的，就从上往下依次查找变量名与key一样的变量，或者变量名忽略大小写后与key一样的变量，第一个匹配的就赋值，后面就算有匹配的也忽略（变量可导出，首字母大写）。
//当接收体中存在json串中匹配不了的项时，解析会自动忽略该项，该项仍保留原值。如变量Work5，保留空值nil。
//json解析后，json串中value，只要是"简单数据"，都会按照默认的类型赋值。
//简单数据：是指不能再进行二次json解析的数据，例如name
//复合数据：是可进行二次甚至多次json解析的，因为它的value也是个可被解析的独立json，例如work1-5。
//对于"复合数据"，如果接收体中配的项被声明为interface{}类型，go都会默认解析成map[string]interface{}类型。如果想直接解析到struct Class对象中，可以将接受体对应的项定义为该struct类型。
//如果不想指定work变量为具体的类型，仍想保留interface{}类型，但又希望该变量可以解析到struct work对象中，可以将该变量定义为json.RawMessage类型。
//被声明为json.RawMessage类型的变量在json解析时，变量值仍保留json的原值，即未被自动解析为map[string]interface{}类型,，可以对该变量进行二次json解析，因为其值仍是个独立且可解析的完整json串,只需再定义一个新的接受体即可。
type User struct {
	Name  string `json:"user_name"`
	Age   int    `json:"age"`
	sex   string
	Work1 *Work1
	Work2 json.RawMessage
	Work3 Work3
	Work4 interface{} `json:"work4"`
	Work5 interface{}
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
	//json字符中的"引号，需用\进行转义，否则编译出错
	data := "{\"user_name\":\"ares\",\"sex\":\"男\",\"age\":18,\"Work1\":{\"work1_name\":\"god1\",\"Salary\":100},\"Work2\":{\"work2_name\":\"god2\",\"Salary\":200},\"Work3\":{\"work3_name\":\"god3\",\"Salary\":300},\"work4\":{\"work4_name\":\"god4\",\"Salary\":400,\"work4_add\":\"cbd\"}}"
	str := []byte(data)
	u1 := User{}
	// Unmarshal的第一个参数是json字符串，第二个参数是接受json解析的数据结构.第二个参数必须是指针，否则无法接收解析的数据，
	err := json.Unmarshal(str, &u1)
	if err != nil {
		fmt.Println("Unmarshal err,", err)
	}
	// {ares 18  0xc0000a41c8 0xc0000a41e0 {god3 300} map[Salary:400 work4_add:cbd work4_name:god4]}  Work2 为*Work2类型
	//  Work2 为json.RawMessage类型  {ares 18  0xc0000a4198 [123 34 119 111 114 107 50 95 110 97 109 101 34 58 34 103 111 100 50 34 44 34 83 97 108 97 114 121 34 58 50 48 48 125] {god3 300} map[Salary:400 work4_add:cbd work4_name:god4] <nil>}

	fmt.Println(u1)
	// 查看类型
	nameType := reflect.TypeOf(u1.Name)
	ageType := reflect.TypeOf(u1.Age)
	sexType := reflect.TypeOf(u1.sex)
	work1Type := reflect.TypeOf(u1.Work1)
	work2Type := reflect.TypeOf(u1.Work2)
	work3Type := reflect.TypeOf(u1.Work3)
	work4Type := reflect.TypeOf(u1.Work4)
	work5Type := reflect.TypeOf(u1.Work5)
	fmt.Println(nameType)  // string
	fmt.Println(ageType)   // int
	fmt.Println(sexType)   // string
	fmt.Println(work1Type) // *main.Work1
	fmt.Println(work2Type) // json.RawMessage
	fmt.Println(work3Type) // main.Work3
	fmt.Println(work4Type) // map[string]interface {}
	fmt.Println(work5Type) // <nil>
}
