package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func(ch chan int) {
		for i := 0; i <= 3; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
	}(ch)

	for {
		if num, ok := <-ch; ok == true {
			fmt.Println("读到数据", num)
		} else {
			break
		}
	}
	//for data := range ch {
	//	fmt.Println("data==", data)
	//	if data == 3 {
	//		break
	//	}
	//}
}
