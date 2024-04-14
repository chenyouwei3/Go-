package main

import "fmt"

func main() {
	// map[KeyType]ValueType
	// KeyType:表示键的类型,只能为基本类型。
	// ValueType:表示键对应的值的类型。
	var m1 map[int]string
	fmt.Println(m1)
	//或者用make定义
	//make(map[KeyType]ValueType, [cap])
	//其中cap表示map的容量，该参数不是必须的
	m2 := make(map[int]string)
	fmt.Println(m2)
	m3 := make(map[int]string, 10)
	fmt.Println(m3)
}
