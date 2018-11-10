package main

import (
    "fmt"
    "time"
    "github.com/samuel/go-zookeeper/zk"
)

func check(e error) {
    if e != nil {
        panic(e)
    }   
}

func main() {
    c, _, err := zk.Connect([]string{"10.149.24.164"}, time.Second) //*10)
    check(err)
    children, stat, ch, err := c.ChildrenW("/") 
    check(err)
    
    fmt.Printf("%+v %+v\n", children, stat)

    content, stat, _ := c.Get("/charge-rtmsg-galaxy/lock")
    fmt.Printf("ip: %s\n", content);

    for {
        select {
            case e := <-ch:
                fmt.Printf("%+v\n", e)
                children, stat, ch, err = c.ChildrenW("/")
                fmt.Printf("%+v %+v\n", children, stat)
        }
        time.Sleep(time.Second * 1)
    }
}
