package main

import (
	"fmt"

	"github.com/toe.royalapp/json-map-struct/utils"
)

type Hello struct {
	Name string
	Age  int
}

func main() {
	person := Hello{
		Name: "Smith",
		Age:  15,
	}
	fmt.Println("map : ", person)

	stm := utils.StructToMap(&person)
	fmt.Println("struct to map : ", stm)

	mts := Hello{}
	_ = utils.MapToStruct(stm, &mts)
	fmt.Println("map to struct : ", mts)

	jsonStr, _ := utils.MapToJson(stm)
	fmt.Println("map to json : ", jsonStr)

	stj, _ := utils.StructToJSON(&person)
	fmt.Println("struct to json : ", stj)

	mapStr, _ := utils.JsonToMap(stj)
	fmt.Println("json to map : ", mapStr)
	fmt.Println("json to map data : ", mapStr["Name"])

	p := Hello{}
	_ = utils.JsonToStruct(stj, &p)
	fmt.Println("json to struct : ", p)
}
