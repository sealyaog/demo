package main

import (
        "fmt"
        "net/url"
        "net/http"
        "io/ioutil"
)

func main() {
        u, _ := url.Parse("http://localhost:8080/xiaoyue")
        q := u.Query()
        q.Set("username", "user")
        q.Set("password", "passwd")
        u.RawQuery = q.Encode()
        res, err := http.Get(u.String());
        if err != nil { 
              fmt.Printf("%s", err) 
              return 
        }
        result, err := ioutil.ReadAll(res.Body) 
        res.Body.Close() 
        if err != nil { 
              fmt.Printf("%s", err)
              return 
        } 
        fmt.Printf("%s\n", result)
}
