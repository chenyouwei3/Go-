package main

import "fmt"

// struct
type Person struct {
	name string
	age  int
}

// NewPerson 构造函数
func NewPerson(name string, age int) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

// 为 person这个类型创建一个学习的方法
func (p Person) Study() {
	fmt.Print("我要卷四各位，或者被各位卷四")
}

// SetAge 设置p的年龄
// 使用指针接收者
func (p *Person) SetAge(newAge int) {
	p.age = newAge
}

//什么时候应该使用指针类型接收者
//需要修改接收者中的值
//
//接收者是拷贝代价比较大的大对象
//
//保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。
//
//注意事项： 非本地类型不能定义方法，也就是说我们不能给别的包的类型定义方法。

func main() {
	p1 := NewPerson("小王子", 25)
	fmt.Println(p1.age) // 25
	p1.SetAge(30)
	fmt.Println(p1.age) // 30
}
