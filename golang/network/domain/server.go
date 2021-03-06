package main

import "net"
import "fmt"

func echoServer(c net.Conn) {
    for {
        buf := make([]byte, 512)
        nr, err := c.Read(buf)
        if err != nil {
            return
        }

        data := buf[0:nr]
        fmt.Printf("Received: %v", string(data))
        _, err = c.Write(data)
        if err != nil {
            //panic("Write: " + err)
        }
    }
}

func main() {
    l, err := net.Listen("unix", "/tmp/echo.sock")
    if err != nil {
        //println("listen error", err.String())
        return
    }

    for {
        fd, err := l.Accept()
        if err != nil {
            //println("accept error", err.String())
            return
        }

        go echoServer(fd)
    }
}
