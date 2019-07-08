package main

import "fmt"
import _ "demo_pkg" // 引入包，包中的init()会执行

func main() {
	fmt.Printf("hello, world\n")
}
