package main

import "fmt"

func sum(x []int, c chan int) {
	sum := 0
	for i, _ := range x {
		sum += i
	}
	c <- sum
}

var done chan bool = make(chan bool)
var msg chan int = make(chan int)

func producer() {
	for i := 0; i < 50; i++ {
		msg <- i
	}
	// close的作用是告诉应用程序已经没有更多的消息，Receiver执行range的时候知道哪里停止
	close(msg) // close 只在Sender侧调用，在Receiver调用会抛出panic
	
	done <- true // 用于阻塞主线程退出，当执行完10次循环，解除阻塞
}

func consumer(s string) {
	for {
		fmt.Printf("consumer %s get %d\n", s, <-msg)
	}
}

func simple(c chan int) {
	c <- 1
	close(c)
}
/*
Channels 等待一端完成，另一端才能获取，即Send完成了，才能Receive, 这样就可以避免显式的锁

Channels Send 或 Receive必须有一个使用goroutine, 否则失败

go 只能作用在函数上
*/

func main() {
	s := []int{1, 2, 3, -9, 10, 20, -6}
	c := make(chan int)
	go sum(s[:len(s) / 2], c)
	go sum(s[len(s) / 2 :], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x + y)

	//c1 := make(chan int)
	go producer()
	go consumer("c1")
	go consumer("c2")

	<-done // 用于阻塞主线程，当获得消息，解除阻塞

	c2 := make(chan int)
	go simple(c2)
	i, ok := <-c2
	if ok {
		fmt.Println(i)
	}
}