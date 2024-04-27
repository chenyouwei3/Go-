package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go func() {
		time.Sleep(time.Second * 1)
		ch <- "脑子进煎鱼了"
	}()

	select {
	case _ = <-ch:
		fmt.Println("其实没有超时")
	case <-time.After(time.Second * 3):
		fmt.Println("煎鱼出去了，超时了！！！")
	}
}
