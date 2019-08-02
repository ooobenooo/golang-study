package main

import (
    "errors"
    "fmt"
    "net/http"
    "net/rpc"
)

type Args struct {
    A, B int
}

type Quotient struct {
    Quo, Rem int
}

type Arith int

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

func main() {
    arith := new(Arith)
    rpc.Register(arith)
    rpc.HandleHTTP()
    
    err := http.ListenAndServe(":9999", nil)
    if err != nil {
        fmt.Println("error occur:", err)
    }
}
