package main

import (
	"fmt"
	"reflect"
)

type valueStructTest0 struct {
	Name   string
	GoodAt []string
}

type valueStructTest1 struct {
	Name   string
	GoodAt []string
}

func main() {
	v1 := valueStructTest0{Name: "煎鱼", GoodAt: []string{"炸", "煎", "蒸"}}
	v2 := valueStructTest1{Name: "煎鱼", GoodAt: []string{"炸", "煎", "蒸"}}
	v3 := valueStructTest1{Name: "煎鱼", GoodAt: []string{"炸", "煎", "蒸"}}
	if reflect.DeepEqual(v3, v2) {
		fmt.Println("相同结构体相同值能通过")
	}
	if reflect.DeepEqual(v1, v2) {
		fmt.Println("不相同结构体相同值能通过")
	}
	fmt.Println("脑子没进煎鱼")
}
