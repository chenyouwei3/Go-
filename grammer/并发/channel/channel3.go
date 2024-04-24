package main

import (
	"fmt"
	"time"
)

// 只写端
func write3(ch chan<- int) {
	ch <- 100
	fmt.Printf("ch addr：%v\n", ch) // 输出内存地址
	ch <- 200
	fmt.Printf("ch addr：%v\n", ch) // 输出内存地址
	ch <- 300                      // 该处数据未读取，后续操作直接阻塞
	fmt.Printf("ch addr：%v\n", ch) // 没有输出
}

// 只读端
func read3(ch <-chan int) {
	// 只读取两个数据
	fmt.Printf("取出的数据data1：%v\n", <-ch) // 100
	fmt.Printf("取出的数据data2：%v\n", <-ch) // 200
}

func main() {

	var ch chan int         // 声明一个双向
	ch = make(chan int, 10) // 初始化

	// 向协程中写入数据
	go write3(ch)

	// 向协程中读取数据
	go read3(ch)

	// 防止主go程提前退出，导致其他协程未完成任务
	time.Sleep(time.Second * 3)
}
