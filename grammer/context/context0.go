package main

import (
	"fmt"
	"sync"
	"time"
)

var wg0 sync.WaitGroup

func worker0() {
	for {
		fmt.Printf("worker->")
		time.Sleep(1 * time.Second)
	}
	wg0.Done()
}

func main() {
	wg0.Add(1)
	go worker0()
	// 如何优雅的实现结束子goroutine
	wg0.Wait()
	fmt.Println("over")
}
