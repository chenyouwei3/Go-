package main

import (
	"fmt"
	"time"
)

// ConcurrentPrint 开启十个线程，每个线程都打印100次 a
func ConcurrentPrint() {
	for i := 0; i < 10; i++ {
		go Print()
	}
}

func Print() {
	for i := 0; i < 100; i++ {
		fmt.Println("a")
	}
}

func main() {
	ConcurrentPrint()
	time.Sleep(5 * time.Second)
}
