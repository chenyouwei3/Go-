package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建100个协程
	for i := 0; i < 100; i++ {
		go func(id int) {
			defer func() {
				if r := recover(); r != nil {
					fmt.Printf("Panic recovered in goroutine %d: %v\n", id, r)
				}
			}()
			if id == 50 {
				// 第50个协程发生panic
				panic("panic in goroutine 50")
			}
			fmt.Printf("Goroutine %d started\n", id)
			time.Sleep(time.Second)
			fmt.Printf("Goroutine %d completed\n", id)
		}(i)
	}
	// 等待一段时间以保证协程有足够的时间执行
	time.Sleep(5 * time.Second)
	fmt.Println("Main goroutine completed")
}
