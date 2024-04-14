package main

import (
	"fmt"
)

// 定义一个类型别名 data，代表 map[int]string 类型
type data map[int]string

func main() {
	// 使用类型别名声明一个变量
	d := data{1: "hello", 2: "world"}

	// 输出类型别名声明的变量
	fmt.Println(d) // 输出：map[1:hello 2:world]
}
