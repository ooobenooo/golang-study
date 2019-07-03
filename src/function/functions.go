package main

import "fmt"

func throwPanic() {
	panic("a panic")
}

func main() {

	a := 2
	b := 3
	n, m := demo(a, b)
	fmt.Printf("a + b = %d || a * b = %d\n", n, m)

	sum := myfunc(1, 2, 3, 4, 5)
	fmt.Printf("1 + 2 + 3 + 4 + 5 = %d\n", sum)

	x := 3
	y := 3

	// 传值
	x1 := add(x)
	fmt.Printf("x1 = %d\n", x1)
	fmt.Printf("x = %d\n", x)

	// 传指针，在变量前加&
	// & 表示获取变量的指针
	// * 表示获取指针的值
	y1 := add1(&y)
	fmt.Printf("y1 = %d\n", y1)
	fmt.Printf("y = %d\n", y)

	z := 3
	z1 := add2(&z)
	fmt.Println(z)
	fmt.Println(z1)  // 打印的内容是z1的内存地址
	fmt.Println(*z1) // 转换为内存地址的值，打印值是6

	slice := []int{1, 2, 3, 4, 5}
	result := check(slice, isOdd)
	fmt.Println(result)

	test()                         // panic出现，recover后跳出函数
	fmt.Println("main func is ok") // 这里继续执行
}

func test() {
	// recover只能和defer一起用
	// 正常状态下recover的值是nil
	// 出现panic,函数终止执行defer声明的函数
	// 不会继续函数
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("catch a panic")
		}
	}()

	throwPanic()          // 出现panic的地方，函数的调用已经终止，和java的try...catch有区别
	fmt.Println("go on?") // 不会打印
}

// go支持函数有多个返回值
// 没有返回值，可以省略第二个括号
// 如果有返回值，必须有return语句
// 命名返回参数的名字，提高可读性
// 命名了返回参数，相当于声明了变量
func demo(a, b int) (sum, mutiple int) {
	return a + b, a * b
}

// 变参，arg是一个slice。
func myfunc(arg ...int) (sum int) {
	// slice 返回2个参数，第一是index,第二个是值
	for t, n := range arg {
		fmt.Println(t)
		sum = sum + n
	}
	return sum
}

// 传值
// a是副本，在函数内发生变化不会影响外部原本变量的值
func add(a int) int {
	a = a + a
	return a
}

// 传指针
// 传递的是a的内存地址，占用8byte,函数修改*a,对a的内存地址直接操作
// 函数返回的类型是int, return返回指针， 返回类型不能声明为*int,否则报错，不能返回指针的指针
// 使用指针的好处是不用复制太多东西，能提升效率，对于大对象的传递建议使用指针
func add1(a *int) int {
	*a = *a + *a
	return *a
}

// 此函数表示返回a的指针，即内存地址
func add2(a *int) *int {
	*a = *a + *a
	return a
}

// 函数可以作为值或参数传递
// 用type定义一个函数类型
type filterFunc func(int) bool

// 定义一个或多个和函数类型一样的入参、返回类型的函数
func isOdd(a int) bool {
	if a%2 == 0 {
		return false
	}

	return true
}

// 定义的函数类型作为函数的一个入参类型
func check(slice []int, f filterFunc) []int {
	var result []int
	for _, v := range slice {
		if f(v) {
			result = append(result, v)
		}
	}

	return result
}
