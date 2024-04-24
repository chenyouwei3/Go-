package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 定义缓冲区大小
const BUFLEN = 5

// 全局位置定义全局变量
var cond *sync.Cond = sync.NewCond(&sync.Mutex{})

// 生产者
func producer(ch chan<- int) {
	for {
		cond.L.Lock()           // 给条件变量对应的互斥锁加锁
		for len(ch) == BUFLEN { // 缓冲区满，则等待消费者消费，这里不能是if
			cond.Wait()
		}
		ch <- rand.Intn(1000) // 写入缓冲区一个随机数
		cond.L.Unlock()       // 生产结束，解锁互斥锁
		cond.Signal()         // 一旦生产后，就唤醒其他被阻塞的消费者
		time.Sleep(time.Second * 2)
	}
}

// 消费者
func consumer(ch <-chan int) {
	for {
		cond.L.Lock()      // 全局条件变量加锁
		for len(ch) == 0 { // 如果缓冲区为空，则等待生产者生产，，这里不能是if
			cond.Wait() // 挂起当前协程，等待条件变量满足，唤醒生产者
		}
		fmt.Println("Receive:", <-ch)
		cond.L.Unlock()
		cond.Signal()
		time.Sleep(time.Second * 1)
	}
}

func main() {

	rand.Seed(time.Now().UnixNano()) // 设置随机数种子

	// 生产消费模型中的
	ch := make(chan int, BUFLEN)

	// 启动10个生产者
	for i := 0; i < 10; i++ {
		go producer(ch)
	}

	// 启动10个消费者
	for i := 0; i < 10; i++ {
		go consumer(ch)
	}

	// 阻塞主程序退出
	for {

	}
}
