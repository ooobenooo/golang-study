package main

import (
	"fmt"
	"runtime"
)

func Say(x string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched() // 停顿片刻，让出CPU资源给其他goroutine使用
		fmt.Println(x)
	}
}

func main() {
	go Say("hello")
	Say("World")
}
