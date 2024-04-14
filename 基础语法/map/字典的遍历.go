package main

import "fmt"

func main() {
	m1 := map[int]string{1: "abc", 2: "def"}
	//迭代遍历1，第一个返回值是key，第二个返回值是value
	for k, v := range m1 {
		fmt.Printf("%d ----> %s\n", k, v)
		//1 ----> abc
		//2 ----> def
	}

	//迭代遍历2，第一个返回值是key，第二个返回值是value（可省略）
	for k := range m1 {
		fmt.Printf("%d ----> %s\n", k, m1[k])
		//1 ----> abc
		//2 ----> def
	}

	//判断某个key所对应的value是否存在, 第一个返回值是value(如果存在的话)
	value, ok := m1[1]
	fmt.Println("value = ", value, ", ok = ", ok) //value =  abc , ok =  true

	value2, ok2 := m1[3]
	fmt.Println("value2 = ", value2, ", ok2 = ", ok2) //value2 =   , ok2 =  false

}
