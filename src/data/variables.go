package main

import "fmt" // 导入包

// golang, 首字母大写的变量和方法是公共的，首字母小写的变量和方法是私有的

// 全局变量只能用var定义，不支持简单声明, 全局变量声明不使用不报错
var globalVariable int = 1 

// const 用于常量声明
const pi = 3.14

// iota 即枚举，默认第一个值0，后面的iota递增
// 由于const分组声明，q也是iota, 递增后实际值为2
const (
	o = iota
	p = iota
	q
)

func main()  {
	// 简短声明，只能在函数内部使用。
	// 局部变量声明必须使用，否则编译错误
	a := 18
	fmt.Println(a)

	// _ 只是一个特殊变量名，没有意义，不能用于操作或计算，声明后不能使用
	_, b := 32, 100
	fmt.Println(b)

	// go支持丰富的数值类型，分32位，64位，分有符号和无符号
	// 不同的类型之间不允许赋值和互相操作
	var c int32 = -10 // 有符号32位 int
	fmt.Println(c)
	var d uint32 = 8 // 无符合32位 int
	fmt.Println(d)
	// var e int32 = c + d 编译错误

	// 支持复数
	var f complex64 = 5+5i
	fmt.Printf("Value is %v\n", f)

	// String 不可变
	var s string = "hello"
	fmt.Printf("string is %s\n", s)

	// String 不可变，但可以将其变为字节数组，改变数组值后再转换为String
	x := []byte(s)
	x[0] = 'c'
	s = string(x)
	fmt.Printf("change string is %s\n", s)

	s1, s2 := "hello", "world"
	fmt.Printf("%s %s\n", s1, s2)
	fmt.Printf(s1 + " " + s2 + "\n") // 用加号链接字符串

	fmt.Printf("q is equal to %v\n", q)

	//数据结构

	// 数组
	// 数组之间的赋值是值的赋值，传入函数的数组是数组的副本，意思在函数中修改了数组，
	// 函数返回后不影响数组内容
	var arr [10]int
	arr[0] = 1
	fmt.Printf("first element of array is %d\n", arr[0])
	fmt.Printf("default element value of array is %d\n", arr[1])
	// arr = [5]int {1,2,3,4,5} 编译错误，arr已经是[10]int，不能改变为[5]int

	//多维数组
	var doubleArr [3][3]int // 不用new就可以初始化一个数组
	doubleArr[0][0] = 1
	doubleArr[1][1] = 2
	fmt.Printf("double array 0,0 value is %d\n", doubleArr[0][0])

	// silce 动态数组，是引用类型，底层是一个数组
	// 使用方式和数组一样，只是不需要声明长度
	var ss []int // ss是引用，没有指向任何内存地址
	ss = arr[:2] // 取arr的0-1的元素创建新数组，ss指向新数组
	fmt.Printf("slice is ref type, base array, slice[0] is %d\n", ss[0])

	ss = []int {2,3,4,5} // slice可以改变数组长度
	fmt.Println(ss)

	var ss1 []int = arr[:2]
	var ss2 []int = arr[:3]

	fmt.Println(arr)

	ss1[0] = 2 // slice是引用，改变了slice元素的值，会对原来数组的数据同时改变
	fmt.Println(ss1) // ss1[0] = 2
	fmt.Println(ss2) // ss2[0] = 2
	fmt.Println(arr) // arr[0] = 2

	// slice的默认容量是10
	// 当cap(slice) - len(slice) == 0, ss1 会指向新的数组，那么修改新slice的内容不会影响原来的数组或slice
	ss1 = append(ss1, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14)
	ss1[0] = 3
	fmt.Println(cap(ss1))
	fmt.Println(cap(ss2))
	fmt.Println(len(ss2))
	fmt.Println(ss1)
	fmt.Println(ss2)
	fmt.Println(arr)

	// map
	// map也是引用类型
	var m map[string]int = make(map[string]int)
	m["ten"] = 10
	fmt.Println(m["ten"])

	var mm map[string]int // 声明一个map
	mm = make(map[string]int) // 初始化，创建一个map对象
	mm["first"] = 100
	fmt.Println(mm["first"])

	v, ok := mm["second"] // key 不存在 ok = false
	if (ok) {
		fmt.Println(v)
	} else {
		fmt.Println("second is not exists.")
	}

	// make 只能用于内置数据结构，channel, slice, map, 返回的是类型，并初始化内部数据结构
	// new 可用于数据类型，返回的是指针
	mmm := make(map[string]int)
	fmt.Println(mmm)

	sss := make([]int, 0, 10)
	fmt.Println(sss)
	fmt.Println(len(sss))
	fmt.Println(cap(sss))

	var ssss []int
	if (ssss == nil) {
		fmt.Println("ssss is nil")
	}

}