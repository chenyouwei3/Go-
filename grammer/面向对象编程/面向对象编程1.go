package main

import "fmt"

// 实现封装
// 首字母大写的字段可以被其他包访问，相当于 public 属性；首字母小写的字段只能在当前包内访问，相当于 private 属性。
type Person1 struct {
	Name string
	age  int
}

// 实现extend
type Person0 struct {
	name string
}

// extend了Person0
type Student0 struct {
	Person0
	grade int
}

func (p Person0) sayHello() {
	fmt.Println("Hello,I am", p.name)
}

func main() {
	s := Student0{Person0{"MIKE"}, 3}
	s.sayHello()
}
