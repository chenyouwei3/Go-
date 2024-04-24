package main

import (
	"fmt"
	"sync"
	"time"
)

// Wait: 该方法会阻塞等待条件变量满足条件。也会对锁进行解锁，一旦收到通知则唤醒，并立即锁定该锁
// Signal: 发送通知(单发)，给一个正在等待在该条件变量上的协程发送通知
// Broadcast: 发送通知(广播），给正在等待该条件变量的所有协程发送通知，容易产生 惊群
func main() {
	cond := sync.NewCond(&sync.Mutex{})
	condition := false
	// 开启一个新的协程，修改变量 condition
	go func() {
		time.Sleep(time.Second * 1)
		cond.L.Lock()
		condition = true // 状态变更，发送通知
		cond.Signal()    // 发信号,锁的权限改变,变成主线程
		cond.L.Unlock()
	}()
	// main协程 是被通知的对象，等待通知
	cond.L.Lock()
	for !condition {
		cond.Wait() // 内部释放了锁（释放后，子协程能拿到锁），并等待通知（消息）
		fmt.Println("获取到了消息")
	}
	cond.L.Unlock() // 接到通知后，会被再次锁住，所以需要在需要的场合释放
	fmt.Println("运行结束")
}
