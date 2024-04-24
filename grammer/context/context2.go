package main

import (
	"context"
	"fmt"
	"time"
)

//context.Background() 用于创建空的父级上下文，作为顶级上下文使用，不包含任何附加的值或取消操作。
//context.TODO() 用于表示一种“无事可做”的状态，用作临时占位符上下文
//一般不应该在实际业务逻辑中使用，而是应该尽快替换为具体的上下文构造函数。

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			// 处理逻辑
			select {
			case <-ctx.Done():
				fmt.Println("over")
				return
			default:
				time.Sleep(time.Second)
				fmt.Println("wait for me")
			}
		}
	}(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	time.Sleep(time.Second * 5)
}
