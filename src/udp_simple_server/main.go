package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	service := ":28080"
	udpAddr, err := net.ResolveUDPAddr("udp", service)
	checkError(err)
	conn, err := net.ListenUDP("udp", udpAddr)
	checkError(err)
	for {
		handleConn(conn) // GO的UDP没有accept， 一个conn共享，不要主动关闭conn
	}
}

func handleConn(conn *net.UDPConn) {
	var buf [512]byte

	_, addr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		return
	}
	dayTime := time.Now().String()
	conn.WriteToUDP([]byte(dayTime), addr)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("error occur", err)
		os.Exit(1)
	}
}
