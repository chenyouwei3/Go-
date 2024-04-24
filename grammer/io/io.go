package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// bufio是在file的基础上封装了一层API，支持更多的功能
func main() {
	//创建
	file, err := os.Open("./test1.go")
	if err != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" + "Does the file exist?\n" + "Have you got access to it?\n")
		return
	}

	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		//该方法的第一个参数 delim 为分隔符，当读取到该字符时，ReadString 方法返回已读取的字符串。第二个返回值 err 表示在读取过程中发生的错误。
		line, err := reader.ReadString('\n') //注意是字符
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		fmt.Print(line)
	}

	//var temp [128]byte
	//n, err := file.Read(temp[:])
	//if err != nil {
	//	log.Printf("err ")
	//	return
	//}
	//fmt.Println(string(temp[:n]))
}
