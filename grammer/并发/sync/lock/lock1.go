// 账户对象，对象内置了金额与锁，对象拥有读取金额、添加金额方法
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// 设定为写模式：与互斥锁使用方式一致，一路只写
	//func (*RWMutex) Lock()				// 锁定写
	//func (*RWMutex) Unlock()			// 解锁写
	//
	//// 设定为读模式：对读执行加锁解锁，即多路只读
	//func (*RWMutex) RLock()
	//func (*RWMutex) RUnlock()
	var rwm sync.RWMutex

	for i := 0; i < 3; i++ {
		go func(i int) {
			fmt.Println("111111111111Try Lock reading i:", i)
			rwm.RLock()
			fmt.Println("222222222Ready Lock reading i:", i)
			time.Sleep(time.Second * 2)
			fmt.Println("33333333Try Unlock reading i: ", i)
			//rwm.RUnlock()
			fmt.Println("4444444444Ready Unlock reading i:", i)

		}(i)
	}

	time.Sleep(time.Millisecond * 100)
	fmt.Println("Try Lock writing ")
	rwm.Lock()
	fmt.Println("Ready Locked writing ")
}
