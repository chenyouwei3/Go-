package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(reflect.TypeOf(new(int)))
	fmt.Println(reflect.TypeOf(new([]int)))
	fmt.Println(reflect.TypeOf(make([]int, 0)))
	fmt.Println(reflect.TypeOf(make(map[int]int)))
}
