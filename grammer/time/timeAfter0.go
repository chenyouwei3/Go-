package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("程序开始运行", time.Now())
	t := time.After(time.Second * 3)
	fmt.Printf("t type =%T\n", t)
	//阻塞三秒
	fmt.Println("开始阻塞", time.Now())
	fmt.Println("t=", <-t)

}
