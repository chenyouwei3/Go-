package main

import (
	"fmt"
	"sync"
)

// CounterCapacity 柜台的容量
// breadCount 面包的数量
var CounterCapacity, breadCount = 20, 0
var l sync.Mutex

var cond = sync.NewCond(&l)

func Eat() {
	for {
		l.Lock()
		//柜台没面包了
		if !(breadCount > 0) {
			// 没面包肯定不能把柜台占着不让面包师放面包吧
			cond.Wait()
		}
		//吃掉一个面包
		breadCount--
		fmt.Printf("吃")
		cond.Signal()
		l.Unlock()
	}
}

func Make() {
	for {
		l.Lock()
		if !(breadCount < CounterCapacity) {
			//柜台满了让人吃
			cond.Wait()
		}
		//将一个面包放到柜台里面
		breadCount++
		fmt.Printf("做")
		cond.Signal()
		l.Unlock()
	}
}

func main() {
	for i := 0; i < 3; i++ {
		go Eat()
		go Make()
	}
	select {}
}
