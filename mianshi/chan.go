package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		for item := range c {
			fmt.Println(item)
		}
	}()
	for i := 0; i <= 10; i++ {
		c <- i
	}
}
