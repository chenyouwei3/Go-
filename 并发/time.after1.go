package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(1 * time.Second)
	defer timer.Stop()
	for {
		select {
		case <-timer.C:
			fmt.Printf("现在是：%d，我脑子进煎鱼了！", time.Now().Unix())
			return // 添加退出循环的条件
		}
	}
}
