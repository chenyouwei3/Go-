package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个定时期，每500毫秒触发一次
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	ticker := time.NewTicker(time.Second * 1)

	// 创建一个定时器，2秒后触发
	stopper := time.NewTimer(time.Second * 5)

	// 声明计数变量
	var i int

	for { // 不断检查通道情况
		select {
		case <-stopper.C: // 计时器到了
			fmt.Println("stop", time.Now().Format("2006-01-02 15:04:05"))
			goto StopHere // 跳出循环
		case <-ticker.C: // 打点触发了
			i++ // 记录触发多少次
			fmt.Println("tick", i, time.Now().Format("2006-01-02 15:04:05"))
		}
	}

StopHere:
	fmt.Println("done")
}
