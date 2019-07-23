package main

import (
    "fmt"
    "strconv"
)

func main() {
    a := strconv.FormatInt(1234, 10) // FormatXXX 都是 xxx转换为string类型
    b, _ := strconv.ParseInt("1233", 10, 64) // ParseXXX 都是 string 转换为 xxx 类型
    c := b + 1
    fmt.Println(a, c)
}