package main

import (
	"fmt"
	"time"
)

// timer 时间到了，只执行一次
// ticker 时间到了，多从执行
// chan<- int 只写单向通道（发送）ch <- 10
// <-chan int 只读单向通道（接受）ch=<-ch
func main() {
	//定时器,时间到了,只执行一次
	var i int
	t := time.NewTimer(time.Second * 2)
	defer t.Stop()
	for {
		//当定时器到达指定的时间时，它会向 t.C 发送一个当前时间，从而通知程序时间已到。
		<-t.C
		fmt.Println("timer running...", i)
		//重置定时器,使其在 2 秒后再次触发
		t.Reset(time.Second * 2)
		i++
	}
}
