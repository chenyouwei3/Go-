package main

import "fmt"

func main() {
	// 数组的每个元素可以通过索引下标来访问，索引下标的范围是从0开始到数组长度减1的位置。
	var a [10]int
	for i := 0; i < 10; i++ {
		a[i] = i + 1
		fmt.Printf("a[%d] = %d\n", i, a[i])
	}
	//range具有两个返回值，第一个返回值是元素的数组下标，第二个返回值是元素的值
	for i, v := range a {
		fmt.Println("a[", i, "]=", v)
	}

}
