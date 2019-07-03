package main

import (
	"fmt"
	"time"
)

func Say(x string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(x)
	}
}

func main() {
	go Say("hello")
	Say("World")
}