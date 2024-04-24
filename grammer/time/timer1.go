package main

import (
	"fmt"
	"time"
)

func main() {
	// timer 时间到了，只执行一次
	t := time.NewTimer(time.Second * 2)
	ch := make(chan bool)
	var i int
	go func(t *time.Timer) {
		select {
		//当定时器到达指定的时间时，它会向 t.C 发送一个当前时间，从而通知程序时间已到。
		case <-t.C:
			fmt.Println("time running...", i)
			//重置定时器,使其在 2 秒后再次触发
			t.Reset(time.Second * 2)
			i++
		case stop := <-ch:
			// 当接收到停止信号，退出 goroutine
			if stop {
				fmt.Println("time stop")
				return
			}
		}
	}(t)
	time.Sleep(10 * time.Second)
	ch <- true
	close(ch)
	time.Sleep(1 * time.Second)
}
