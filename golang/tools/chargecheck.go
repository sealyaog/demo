package main

import (
    "golang.org/x/crypto/ssh"
    "os"
    "bufio"
    "fmt"
)

type MatchInfo map[string] string
type DBInfo map[string] DBSchema
type ConnInfo map[int] ConnSchema 

type DebugSchema struct {
    MachineTable MatchInfo 
    DBTable DBInfo 
}

type DBSchema struct {
    user string 
    passwd string
    shardnum int
    partnum int 
    ConnTable ConnInfo 
}

type ConnSchema struct {
    dbhost string
    dbport int
    instname string
}

var GInfo = map[string] DebugSchema {
   "xuri": DebugSchema {
        MatchInfo {
            "charge-route": "xxxxx",            
            "rtchage":"yyyy",
            "streamcharge":"zzzz",
        },

        DBInfo { 
            "tran": DBSchema{
                "for_dev", 
                "sogouAD!@#",
                2,
                8,
                ConnInfo{
                    1: ConnSchema {
                        "10.139.19.47",
                        3401,
                        "charge01",
                    },
                    2: ConnSchema {
                        "10.139.19.48 ",
                        3402,
                        "charge02", 
                    },
                },
            },
        },
    },
    "dubhe": DebugSchema{
        MatchInfo {
            "charge-route": "xxxxx",    
            "rtchage":"yyyy",
            "streamcharge":"zzzz",
        }, 
        DBInfo { 
            "tran": DBSchema{
                "for_dev", 
                "sogouAD!@#",
                2,  
                8,  
                ConnInfo{
                    1: ConnSchema {
                        "10.139.19.47",
                        3401,
                        "charge01",
                    },  
                    2: ConnSchema {
                        "10.139.19.48 ",
                        3402,
                        "charge02", 
                    },  
                },  
            },  
        },        
    },
    "chenxing": DebugSchema{
        MatchInfo {
            "charge-route": "xxxxx",    
            "rtchage":"yyyy",
            "streamcharge":"zzzz",
        },  
        DBInfo {
            "tran": DBSchema{
                "for_dev",
                "sogouAD!@#",
                2,
                8,
                ConnInfo{
                    1: ConnSchema {
                        "10.139.19.47",
                        3401,
                        "charge01",
                    },
                    2: ConnSchema {
                        "10.139.19.48 ",
                        3402,
                        "charge02",
                    },
                },
            },
        },
    },
    "haoyue": DebugSchema{
        MatchInfo {
            "charge-route": "xxxxx",    
            "rtchage":"yyyy",
            "streamcharge":"zzzz",
        },  
        DBInfo {
            "tran": DBSchema{
                "for_dev",
                "sogouAD!@#",
                2,
                8,
                ConnInfo{
                    1: ConnSchema {
                        "10.139.19.47",
                        3401,
                        "charge01",
                    },
                    2: ConnSchema {
                        "10.139.19.48 ",
                        3402,
                        "charge02",
                    },
                },
            },
        },
    },
    "yinhe": DebugSchema{
        MatchInfo {
            "charge-route": "xxxxx",    
            "rtchage":"yyyy",
            "streamcharge":"zzzz",
        },  
        DBInfo {
            "tran": DBSchema{
                "for_dev",
                "sogouAD!@#",
                2,
                8,
                ConnInfo{
                    1: ConnSchema {
                        "10.139.19.47",
                        3401,
                        "charge01",
                    },
                    2: ConnSchema {
                        "10.139.19.48 ",
                        3402,
                        "charge02",
                    },
                },
            },
        },
    },
};


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
