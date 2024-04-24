package main

import "fmt"

// 定义一个接口
type Foo interface {
	FooMethod()
}

// 定义一个结构体
type MyStruct struct{}

// 实现接口中的方法
func (ms MyStruct) FooMethod() {
	fmt.Println("FooMethod called")
}

func main() {
	// 创建一个 MyStruct 的实例
	myStruct := MyStruct{}

	// 判断 MyStruct 是否实现了 Foo 接口
	_, ok := interface{}(myStruct).(Foo)
	if ok {
		fmt.Println("MyStruct 实现了 Foo 接口")
	} else {
		fmt.Println("MyStruct 没有实现 Foo 接口")
	}
}
