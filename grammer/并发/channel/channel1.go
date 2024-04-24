package main

import (
	"fmt"
	"time"
)

// 写端
func write1(ch chan int) {
	ch <- 100
	fmt.Printf("ch addr：%v\n", ch) // 输出内存地址
	ch <- 200
	fmt.Printf("ch addr：%v\n", ch) // 输出内存地址
	ch <- 300                      // 写入第三个，造成阻塞
	fmt.Printf("ch addr：%v\n", ch) // 没有输出
}
func main() {

	var ch chan int        // 声明一个有缓冲的channel
	ch = make(chan int, 2) //ch = make(chan int, 3)// 可以写入2个数据

	// 向协程中写入数据
	go write1(ch)

	// 防止主go程提前退出，导致其他协程未完成任务
	time.Sleep(time.Second * 3)
}
