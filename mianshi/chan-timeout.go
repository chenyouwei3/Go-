package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个通道用于接收结果
	result := make(chan string)

	// 模拟一个耗时的操作，假设需要 2 秒完成
	go func() {
		time.Sleep(2 * time.Second)
		result <- "Operation completed"
	}()

	// 使用 select 语句等待通道操作完成或超时
	select {
	case res := <-result:
		// 如果结果已经准备好，则从通道中接收并打印结果
		fmt.Println(res)
	case <-time.After(5 * time.Second):
		// 如果超过 1 秒仍然没有接收到结果，则执行超时操作
		fmt.Println("Timeout")
	}
}
