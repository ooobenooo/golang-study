package main

import (
	"fmt"
	"os"
	"net/rpc"
	//"net/rpc/jsonrpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:9999")  // gob编码
	// client, err := jsonrpc.Dial("tcp", "127.0.0.1:9999") // json编码
	if err != nil {
		fmt.Println("error occur:", err)
		os.Exit(1)
	}

	arg := Args{17 ,8}
	var reply int
	err = client.Call("Arith.Multiply", arg, &reply) // 必须是&reply，否则在方法中返回不会变
	if err != nil {
		fmt.Println("error occur:", err)
	}
	fmt.Println("17 * 8 =", reply)

	name := "Ben"
	var resp string
	err = client.Call("Hello.SayHi", name, &resp)
	if err != nil {
		fmt.Println("error occur:", err)
	}
	fmt.Println(resp)
}
