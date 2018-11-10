package main

import (
    "golang.org/x/crypto/ssh"
    "os"
    "bufio"
    "fmt"
)

//cmd example: date

func check(e error) {
    if e != nil {
        panic(e)
    }   
}

func SSH(user, password, ip_port string) {
    PassWd := []ssh.AuthMethod{ssh.Password(password)}
    Conf := ssh.ClientConfig{User: user, Auth: PassWd}
    Client, _ := ssh.Dial("tcp", ip_port, &Conf)
    defer Client.Close()
    a := bufio.NewReader(os.Stdin)
    for {
        b, _, z := a.ReadLine()
        if z != nil {
            return
        }
        command := string(b)
        
        session, err := Client.NewSession(); 
        check(err)
        
        defer session.Close()
        //session.Stdout = os.Stdout
        //session.Stderr = os.Stderr
        result,err := session.Output(command)
        check(err)
        fmt.Println("result: %s",string(result));
    }
}

func main() {
    SSH("for_dev", "(jmZ%znqEc", "10.136.20.97:22")
}
