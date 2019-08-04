package main

import (
    "errors"
    "fmt"
    "net"
	"net/rpc"
	// "net/rpc/jsonrpc"
	"os"
)

type Args struct {
    A, B int
}

type Quotient struct {
    Quo, Rem int
}

type Arith int

// 方法必须是导出，返回值必须是指针
func (a *Arith) Multiply(arg *Args, reply *int) error {
    *reply = arg.A * arg.B
    return nil
}

func (a *Arith) Divide(arg *Args, quo *Quotient) error {
    if arg.B == 0 {
        return errors.New("divide by zero")
    }
    
    quo.Quo = arg.A / arg.B
    quo.Rem = arg.A % arg.B
    return nil
}

type Hello string

func (h *Hello) SayHi(name *string, reply *string) error {
	*reply = "Hello " + *name
	return nil
}

func main() {
	arith := new(Arith) 
	rpc.Register(arith) // 注册

	hello := new(Hello)
	rpc.Register(hello) // 注册

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":9999")
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept() // 阻塞式的
		if err != nil {
			continue
		}
		go rpc.ServeConn(conn) // 利用goroutine处理多个请求 gob 编码
		// go jsonrpc.ServeConn(conn) // 用json方式编码
	}
	
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
