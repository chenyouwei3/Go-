package main

func main() {
	ch := make(chan int) // 创建一个无缓冲 channel
	close(ch)            // 关闭 channel
	ch <- 42             // 尝试向已关闭的 channel 发送数据，导致 panic
}
