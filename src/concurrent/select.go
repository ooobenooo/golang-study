package main

import (
	"fmt"
)

// 当有多个Channel时，用select, 当case条件满足就执行语句，如果多个case都成立, 则随机调用。
func fib(c chan int, quit chan bool) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x + y
		case <- quit:
			fmt.Println("quit")
			return
		// default: // 每次都会执行，无阻塞的，不管哪个case有效
		//	fmt.Println("OK")
		}
	}
}

func main() {

	c := make(chan int)
	quit := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- true
	}()

	fib(c, quit)
}