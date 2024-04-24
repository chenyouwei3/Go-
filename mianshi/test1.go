package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cannel := context.WithCancel(context.Background())
	defer cannel()
	ch := make(chan int)
	go func() {
		defer close(ch)
		for {
			select {
			case <-ctx.Done():
				fmt.Println("收到关闭通知")
			case ch <- 1:
				fmt.Println("向通道发送数据")
				time.Sleep(time.Second)
			}
		}
	}()
	time.Sleep(3 * time.Second) // 等待一段时间
	cannel()                    // 发出取消通知
	// 等待goroutine退出
	time.Sleep(time.Second)
	fmt.Println("主函数退出")
}
