package main

import (
	"fmt"
	"time"
)

// 命令行会不断地输出 tick，同时可以使用 fmt.Scanln() 接受用户输入。
// 两个环节可以同时进行，直到按 Enter键时将输入的内容写入 input变量中并返回， 整个程序终止。
func running() {
	var times int
	for {
		times++
		fmt.Println("tick:", times)
		time.Sleep(time.Second)
	}
}

func main() {
	go running()
	var input string
	fmt.Scanln(&input)
}
