package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go cout("sheep", c)
	for msg := range c {
		fmt.Println(msg)
	}
}

func cout(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing
		time.Sleep(time.Second * 1)
	}
	close(c)
}
