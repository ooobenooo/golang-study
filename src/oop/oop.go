package main

import (
	"fmt"
)

// struct是一种类型，Rectangle是自定义类型。
type Rectangle struct {
	length int
	width  int
}

// 任意的内置类型都可以衍生出自定义类型
// 自定义类型是内置类型的别名
type age int

type monthes []string

type kv map[string]int

type some struct {
	age
	monthes
}

// 类型的方法 (r Rectangle) 是接受者
// Receiver是值类型，即是一个拷贝，也可以是指针类型
func (r Rectangle) area() int {
	return r.length * r.width
}

// 如果使用指针，method内部不强制用*做转换，调用时也不强制&转换，go内部帮转
func (r *Rectangle) SetLength(l int) {
	r.length = l
}

type person struct {
	name  string
	age   int
	phone string
}

// struct 可以继承
// 如果有相同字段时，最外层的为准
type student struct {
	person     // 匿名字段
	speciality string
	phone      string
}

type teacher struct {
	person
	subject string
}

// 方法可以重写
func (t teacher) SayHi() {
	fmt.Printf("My name is %s, my subject is %s\n", t.name, t.subject)
}

// method 可以继承
// student 内部有person匿名字段，person实现了SayHi(), student也有SayHi()
func (p person) SayHi() {
	fmt.Printf("My name is %s\n", p.name)
}

func main() {
	// struct 的初始化用{}赋值
	// 构造函数方式需要把所有参数列出
	ben := person{"ben", 30, "12345678"}
	// 不同的赋值方式
	// field:value 不需要初始化所有字段，按需初始
	isaac := student{person: person{age: 6, name: "Isaac"}, speciality: "financial"}
	bob := student{person{"bob", 20, "33333333"}, "computer science", "87654321"}
	fmt.Println(ben)
	fmt.Println(isaac)
	fmt.Println(bob)
	fmt.Println(bob.phone) // 87654321

	r := Rectangle{10, 15}
	area := r.area()
	fmt.Println(area)

	// r不需要写成&r,当然写&r也支持
	r.SetLength(20)
	fmt.Println(r.area())

	s := some{10, monthes{"JAN", "FEB"}}
	fmt.Printf("age of s is %d\n", s.age)
	fmt.Println(s.monthes)

	ben.SayHi()
	isaac.SayHi()

	t := teacher{person{"ann", 20, "9999999"}, "IT"}
	t.SayHi()
}
