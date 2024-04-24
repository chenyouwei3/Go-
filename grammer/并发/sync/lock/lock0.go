package main

import (
	"fmt"
	"sync"

	//"sync"
	"time"
)

func main() {

	var mutex sync.Mutex

	num := 0

	// 开启10个协程，每个协程都让共享数据 num + 1
	for i := 0; i < 1000; i++ {

		go func() {
			mutex.Lock() // 加锁，阻塞其他协程获取锁
			num += 1
			mutex.Unlock() // 解锁
		}()
	}

	// 大致模拟协程结束 等待5秒
	time.Sleep(time.Second * 5)

	// 输出1000，如果没有加锁，则输出的数据很大可能不是1000
	fmt.Println("num = ", num)

}
