package main

import (
    "fmt"
    "os"
)

func main() {

    fmt.Println("test the os package")
    fmt.Println("Args len", len(os.Args))
    for i, arg := range os.Args {
        fmt.Printf("arg[%v] : %v\n", i, arg);
    }

    _, err := os.Stat(os.Args[0])
    if err != nil {
        fmt.Println(err)
    }
}
