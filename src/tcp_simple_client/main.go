package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":18080")
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	defer conn.Close()
	checkError(err)
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)
	reply := make([]byte, 1024)
	n, err := conn.Read(reply)
	checkError(err)
	fmt.Println(string(reply[0:n]))
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("error occur:", err)
		os.Exit(1)
	}
}
