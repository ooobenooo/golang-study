package main

import (
	"fmt"
)

/*
Buffered-Channels 是不需要一定用go执行Send或Receive的
*/


func main() {

	c := make(chan int, 2)
	// fmt.Println(<-c) 没有数据马上Receiver会报错
	c <- 1
	c <- 2
	// 如果在这里增加 c <- 3 会报错，因为Buffered Channel已经满了，继续Send会异常
	fmt.Println(<-c)
	fmt.Println(<-c)
	c <- 3
	fmt.Println(<-c) // 如果没有上一行的 c <- 3直接执行此行，也会报错，缓冲区为空会Receive不了数据。
}