package main

import (
	"fmt"
	"log"
	"reflect"

	"github.com/tealeg/xlsx"
)

type Username struct {
	Name string `excel:"姓名"`
	Age  int    `excel:"年龄"`
}

func main() {
	var data = []Username{
		{
			Name: "张三",
			Age:  13,
		},
		{
			Name: "李四",
			Age:  14,
		},
		{
			Name: "王五",
			Age:  15,
		},
	}
	excel, err := Struct2Xlsx(data)
	if err != nil {
		log.Fatal(err)
	}
	if err := excel.Save("username.xlsx"); err != nil {
		log.Fatal(err)
	}
}

func getStructTagList(v interface{}, tag string) []string {
	var resList []string
	if v == nil {
		return resList
	}
	var item interface{}
	switch reflect.TypeOf(v).Kind() {
	case reflect.Slice, reflect.Array:
		values := reflect.ValueOf(v)
		if values.Len() == 0 {
			return resList
		}
		item = values.Index(0).Interface()
	case reflect.Struct:
		item = reflect.ValueOf(v).Interface()
	default:
		panic(fmt.Sprintf("type %v not support", reflect.TypeOf(v).Kind()))
	}
	typeOf := reflect.TypeOf(item)
	fieldNum := typeOf.NumField()
	for i := 0; i < fieldNum; i++ {
		resList = append(resList, typeOf.Field(i).Tag.Get(tag))
	}
	return resList
}

func getTagValMap(v interface{}, tag string) map[string]string {
	resMap := make(map[string]string)
	if v == nil {
		return resMap
	}
	typeOf := reflect.TypeOf(v)
	fieldNum := typeOf.NumField()
	for i := 0; i < fieldNum; i++ {
		structField := typeOf.Field(i)
		tagValue := structField.Tag.Get(tag)
		val := reflect.ValueOf(v).FieldByName(structField.Name)
		resMap[tagValue] = fmt.Sprintf("%v", val.Interface())
	}
	return resMap
}

func struct2MapTagList(v interface{}, tag string) []map[string]string {
	var resList []map[string]string
	switch reflect.TypeOf(v).Kind() {
	case reflect.Slice, reflect.Array:
		values := reflect.ValueOf(v)
		for i := 0; i < values.Len(); i++ {
			resList = append(resList, getTagValMap(values.Index(i).Interface(), tag))
		}
		break
	case reflect.Struct:
		val := reflect.ValueOf(v).Interface()
		resList = append(resList, getTagValMap(val, tag))
		break
	default:
		panic(fmt.Sprintf("type %v not support", reflect.TypeOf(v).Kind()))
	}
	return resList
}

func Struct2Xlsx(v interface{}) (*xlsx.File, error) {
	var tag = "excel"
	tagList := getStructTagList(v, tag)
	mapTagList := struct2MapTagList(v, tag)
	excelFile := xlsx.NewFile()
	sheet, err := excelFile.AddSheet("Sheet1")
	if err != nil {
		return nil, err
	}
	headerRow := sheet.AddRow()
	for _, tagVal := range tagList {
		headerRow.SetHeightCM(0.5)
		headerRow.AddCell().Value = tagVal
	}
	for _, mapTagVal := range mapTagList {
		row := sheet.AddRow()
		for _, tagVal := range tagList {
			row.SetHeightCM(0.5)
			row.AddCell().Value = mapTagVal[tagVal]
		}
	}
	return excelFile, nil
}
