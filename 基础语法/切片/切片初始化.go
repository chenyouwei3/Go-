package main

import "fmt"

func main() {
	//初始化一个空的切片
	var a1 []int //就比数组少个中括号里的长度
	fmt.Println(a1)
	a2 := []int{}
	fmt.Println(a2)
	//或者用make函数
	//make([]T, length, capacity) //capacity省略，则和length的值相同
	var a3 []int = make([]int, 0)
	fmt.Println(a3)
	a4 := make([]int, 0, 0)
	fmt.Println(a4)
	a5 := []int{1, 2, 3} //创建切片并初始化值
	fmt.Println(a5)
}
