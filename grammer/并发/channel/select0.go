package main

import (
	"fmt"
	"time"
)

func fn1(ch chan string) {
	time.Sleep(time.Second * 3)
	ch <- "fn1111"
}

func fn2(ch chan string) {
	time.Sleep(time.Second * 2)
	ch <- "fn2222"
}

func main() {
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(1e9) // 等待1秒钟
		timeout <- true
	}()

	ch1 := make(chan string)
	go fn1(ch1)

	ch2 := make(chan string)
	go fn2(ch2)

	select {
	case r1 := <-ch1:
		fmt.Println("r1=", r1)
	case r2 := <-ch2:
		fmt.Println("r2=", r2)
	case <-timeout: // 没有从-ch中取到数据，此时能从timeout中取得数据
	}
}
