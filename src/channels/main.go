package main

import "fmt"

func sum(x []int, c chan int) {
	sum := 0
	for _, v := range x {
		sum += v
	}
	c <- sum // 将sum发送到c
}

func main() {
	s := []int{1, 2, 3, -9, 10, 20}
	c := make(chan int)
	go sum(s[:len(s)/2], c) // slice 的后半部分累加
	go sum(s[len(s)/2:], c) // slice 的前半部分累加
	x, y := <-c, <-c        // 将c赋值给x,ydd
	fmt.Println(x, y, x+y)
}
