package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	service := ":18080"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept() // 接收请求
		if err != nil {
			fmt.Println("accept connection error,", err)
			continue // 接收请求出现错误不结束服务器，继续循环
		}
		go handleClient(conn) // 多线程处理客户端请求
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	dayTime := time.Now().String()
	conn.Write([]byte(dayTime))
}

func checkError(err error) {
	if err != nil {
		fmt.Println("error occur:", err)
		os.Exit(1)
	}
}
