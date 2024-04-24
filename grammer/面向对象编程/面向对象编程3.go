package main

import "fmt"

// 定义形状接口
type Shape interface {
	Area() float64
}

// 定义矩形结构体
type Rectangle struct {
	width, height float64
}

// 实现形状接口的 Area 方法
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// 定义圆形结构体
type Circle struct {
	radius float64
}

// 实现形状接口的 Area 方法
func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func main() {
	// 创建一个矩形对象
	rect := Rectangle{width: 5, height: 3}
	// 创建一个圆形对象
	circle := Circle{radius: 2}

	// 声明一个形状接口类型的变量
	var shape Shape

	// 将矩形对象赋值给形状变量
	shape = rect
	// 调用形状对象的 Area 方法，实现多态
	area := shape.Area()
	fmt.Println("Rectangle Area:", area)

	// 将圆形对象赋值给形状变量
	shape = circle
	// 调用形状对象的 Area 方法，实现多态
	area = shape.Area()
	fmt.Println("Circle Area:", area)
}
