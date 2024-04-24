package main

import (
	"fmt"
	"time"
)

// 生产者
func producer(ch chan<- int) {
	i := 1
	for {
		ch <- i
		fmt.Println("Send:", i)
		i++
		time.Sleep(time.Second * 1) // 避免数据流动过快
	}
}

// 消费者
func consumer(ch <-chan int) {
	for {
		v := <-ch
		fmt.Println("Receive:", v)
		time.Sleep(time.Second * 2) // 避免数据流动过快
	}
}

func main() {

	// 生产消费模型中的缓冲区
	ch := make(chan int, 5)
	// 启动生产者
	go producer(ch)
	// 启动消费者
	go consumer(ch)
	// 阻塞主程序退出
	for {

	}
}
