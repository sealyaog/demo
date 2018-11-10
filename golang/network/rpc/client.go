package main

import (
    "fmt"
    "net/rpc"
)

type Args struct {
    A, B int 
}

func main () {
    client, err := rpc.DialHTTP("tcp", "127.0.0.1" + ":1234")
    if err != nil {
        return  
    }
         
    args := &Args{7,8}
    var reply int 
    err = client.Call("Arith.Multiply", args, &reply)
    if err != nil {
        return
    }
    fmt.Printf("Result: %d*%d=%d\n", args.A, args.B, reply) 
}
