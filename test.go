package main

import (
	"fmt"
)

// 定义一个自定义类型 MyInt，基于内置的 int 类型
type MyInt int

// 为自定义类型 MyInt 添加一个方法 Double，用于将自身乘以2并返回结果
func (m MyInt) Double() MyInt {
	return m * 2
}

func main() {
	// 声明一个 MyInt 类型的变量
	num := MyInt(5)

	// 调用自定义类型的方法
	fmt.Println(num.Double()) // 输出：10
}
