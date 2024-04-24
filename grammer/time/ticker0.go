package main

import (
	"fmt"
	"time"
)

func main() {
	//定时期,时间到了,一直执行
	t := time.NewTicker(time.Second * 2)
	defer t.Stop()
	for {
		//当定时期到达指定的时间时，它会向 t.C 发送一个当前时间，从而通知程序时间已到。
		<-t.C
		fmt.Println("Ticker running...")
	}
}
