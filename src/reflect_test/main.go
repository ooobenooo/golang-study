package main

import (
    "fmt"
    "reflect"
)

type Person struct {
    name string
    age int
}

func main() {
    var i float64 = 4.30
    t := reflect.TypeOf(i) // 获取类型
    fmt.Println(t) // 输出 float64
    
    v := reflect.ValueOf(i) // 基本类型不需要用指针， 不用指针是不能反射修改值的
    fmt.Println(v.String())
    fmt.Println(v.Float())
    
    vp := reflect.ValueOf(&i) // 用指针
    vp.Elem().SetFloat(5.0) // 必须用Elem() 才能修改值，Elem()方法能获取指针的值
    fmt.Println(i)
    
    p := Person {"ben", 20}
    pr := reflect.ValueOf(&p) // 反射必须是公共类型, 私有类型（首字母小写）反射失败， struct 类型必须用指针
    fieldName := pr.Elem().Type().Field(0).Name
    fmt.Println(fieldName)
}