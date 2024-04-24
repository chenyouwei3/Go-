package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var money = 10000
	var m sync.Mutex
	// 开启10个协程，每个协程内部 循环1000次，每次循环值+10
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			m.Lock()
			defer wg.Done()
			for j := 0; j < 100; j++ {
				money += 10 //  多个协程对 money产生了竞争
			}
			m.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("最终的monet = ", money) // 应该输出20000才正确，但是每次运行输出结果不正确

}
