package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	values := []string{"a", "b", "c"}
	for _, v := range values {
		wg.Add(1)
		//go func(test string) {
		//	fmt.Println(test)
		//	wg.Done()
		//}(v)
		go func() {
			fmt.Println(v)
			wg.Done()
		}()
	}
	wg.Wait()
}
