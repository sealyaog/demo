package main

import (
    "net/rpc"
    "net/http"
    "log"
    "net"
)
 
 
type Args struct {
    A, B int
}
 
type Arith int
 
func (t *Arith) Multiply(args *Args, reply *int) error {
    *reply = args.A * args.B 
    return nil
}
 
func main() {
    arith := new(Arith)
 
    rpc.Register(arith)
    rpc.HandleHTTP()
        
    l, e := net.Listen("tcp", ":1234")
    if e != nil {
        log.Fatal("listen error:", e)
    }
    http.Serve(l, nil)
}

