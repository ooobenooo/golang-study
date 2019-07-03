package main

import "fmt"

func main() {
	ok := true
	if ok { // 必须有花括号
		fmt.Println("ok")
	}

	// if 语句可以声明变量，并在if语句块中有效
	if x := 10; x > 5 {
		fmt.Printf("ok %d\n", x)
	} else {
		fmt.Println("false")
	}

	demoGoTo()
	fmt.Println()
	demoFor()
	demoSwitch(1)
}

func demoGoTo() {
	i := 0
Here: // 标签，类似html中的锚点
	fmt.Printf("i = %d\n", i)
	i++
	if i < 10 {
		goto Here
	}
}

// for 循环可以当 for 和 while用
// for expression1; expression2; expression3 {}
// 当for用，expression1, expression2, expression3都必须的
// 当while用，只要expression2
// for 可结合break, continues使用
func demoFor() {

	fmt.Println("for loop")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("while loop")
	count := 0
	for count < 10 {
		fmt.Println(count)
		count++
	}
}

// go的switch的case不需要break
func demoSwitch(a int) {
	switch a {
	case 1:
		fmt.Println("first")
	case 2:
		fmt.Println("second")
	default:
		fmt.Println("zero")
	}
}
