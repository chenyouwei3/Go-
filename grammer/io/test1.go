package main

import (
	"fmt"
	"reflect"
)

func main() {

	string1 := "chenyouwei3"
	fmt.Println(string1[:])
	fmt.Println(reflect.TypeOf(string1[:]))
}
