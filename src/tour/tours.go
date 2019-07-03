// 所有GO程序都有mainb包
package main

// import只要换行就是新的导入
import (
	"fmt"
	"math"
	"math/rand"
)

func add(x int, y int) int {
	return x + y
}

func minus(x, y int) int {
	return x - y
}

func swap(x string, y string) (string, string) {
	return y, x
}

/*
return 变量名在定义方法时已经声明，这种方式适合短方法。
*/
func namedReturnFunc(sum int) (x, y int) {
	x = sum * 4
	y = sum - 5
	return
}

func main() {
	fmt.Println("try rand int ", rand.Intn(100))
	fmt.Printf("try math sqrt %g.\n", math.Sqrt(7)) // Printf 格式化
	fmt.Printf("try math.Pi %g\n", math.Pi) // 字段名第一个字母大小为公开访问字段，相当于Java的public, 小写相当与Java的private
	fmt.Println("try func", add(5, 6))
	fmt.Println(minus(6, 5))
	fmt.Println(swap("Hello", "Ben"))
	a, b := swap("dd", "cc") // := 左边的变量之前没有使用过，用这种方法可以创建变量并赋值，不需要用var
	fmt.Println(a, b)
	e, d := namedReturnFunc(100)
	fmt.Println(e, d)

	// 和scala类似，用 var 声明变量
	var c, java, scala bool
	var i int
	fmt.Println(c, java, scala, i)

	// Go支持类型推断，声明变量时不需要指定类型
	var x, y, z = "x", "y", 1
	fmt.Println(x, y, z)

	// Go的基础类型非常丰富
	// bool default value false
	// string default value ""
	// int  int8  int16  int32  int64 default value 0
	// uint uint8 uint16 uint32 uint64 uintptr default value 0
	// byte // alias for uint8 default value 0
	// rune // alias for int32 default value 0
    // represents a Unicode code point
	// float32 float64 default value 0
	// complex64 complex128 default value 0

	// Go的类型转换需要显式处理
	var q int = 100
	var f float64 = float64(q)
	fmt.Println(f)

	// 常量 const, 不能用 := 声明常量
	const g = 1
	// g = 2 编译错误
	fmt.Println(g)


}