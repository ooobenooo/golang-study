package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Servers struct {
	XMLName     xml.Name `xml:"servers"`     // ``中的内容是struct tag, 是struct的一个特性，基于这个做反射
	Version     string   `xml:"version,arr"` // 字段名称必须是公共访问级别
	Srvs        []Server `xml:"server"`
	Description string   `xml:",innerxml"`
}

type Server struct {
	ServerName string `xml:"serverName"`
	ServerIP   string `xml:"serverIP"`
}

func main() {
	file, err := os.Open("server.xml")
	if err != nil {
		fmt.Println("open file error", err)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("read data error", err)
		return
	}

	v := Servers{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Println("parse xml to object error.", err)
		return
	}

	fmt.Println(v)

	out := &Servers{Version: "1.0"}
	out.Srvs = append(out.Srvs, Server{"shenzhen", "8.8.8.8"})
	out.Srvs = append(out.Srvs, Server{"guangzhou", "4.4.4.4"})

	output, err := xml.MarshalIndent(out, " ", "    ") // 序列化对象
	if err != nil {
		fmt.Println("serialize object error", err)
		return
	}

	os.Stdout.Write([]byte(xml.Header)) // 写入头
	os.Stdout.Write(output)
}
