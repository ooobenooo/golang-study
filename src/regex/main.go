package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "I am learning Go language"
	re, _ := regexp.Compile("[a-z]{2,4}")

	fmt.Println("Find All")
	s := re.FindAll([]byte(str), -1)
	for _, v := range s {
		fmt.Println(string(v))
	}

	// 找出匹配的子字符串的开始和结束位置
	i := re.FindAllStringIndex(str, -1)
	fmt.Println(i)

	src := []byte(`
		call hello alice
		hello bob
		call hello eve
	`)

	// P<> 可以将匹配的转换为一个参数，并声明参数变量名
	pat := regexp.MustCompile(`(?m)(call)\s+(?P<cmd>\w+)\s+(?P<arg>.+)\s*$`)
	var res []byte
	for _, s := range pat.FindAllSubmatchIndex(src, -1) {
		// Expand通过模板，填入变量值
		res = pat.Expand(res, []byte("$cmd('$arg')\n"), src, s)
	}
	fmt.Println(string(res))
}
