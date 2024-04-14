package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//i是下标，v是元素的值
	for i, v := range slice {
		fmt.Printf("切片[%d]=%d, ", i, v)
		//切片[0]=-1, 切片[1]=1, 切片[2]=2, 切片[3]=3, 切片[4]=4, 切片[5]=5, 切片[6]=6, 切片[7]=7, 切片[8]=8, 切片[9]=9,
	}
	fmt.Println("\n")
}
