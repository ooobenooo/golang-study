package main

import (
	"encoding/json"
	"fmt"
)

// 字段必须是可导出的，根据字段名称映射。
// 只会解析字段名的值，其他忽略
type Person struct {
	Name string `json:"name"` // 利用tag控制json输出中字段的首字母大小写
	Age  int    `json:"age"`
}

// 字段必须是可导出的
type PersonSlice struct {
	Persons []Person `json:"persons"`
}

func main() {
	jsonStr := `{"persons":[{"name":"ben","age":38,"weight":120},{"name":"peter","age":25}]}`

	var p PersonSlice
	err := json.Unmarshal([]byte(jsonStr), &p)
	if err != nil {
		fmt.Println("parse json error.", err)
	}

	fmt.Println(p)

	// 当不知道结构体时，可以用interface{}
	// 解析是一个map对象
	var i interface{}
	err = json.Unmarshal([]byte(jsonStr), &i)
	if err != nil {
		fmt.Println("parse json error,", err)
	}

	m := i.(map[string]interface{})
	fmt.Println(m)

	var output PersonSlice
	output.Persons = append(output.Persons, Person{"ken", 40})
	output.Persons = append(output.Persons, Person{"Anne", 20})

	b, err := json.Marshal(output)
	if err != nil {
		fmt.Println("parse json error,", err)
	}

	fmt.Println(string(b))
}
