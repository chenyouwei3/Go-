package main

import "fmt"

func main() {
	array := [5]int{0, 1, 2, 3, 4}
	for i, v := range array {
		fmt.Println(i, "is", v)
	}
}
