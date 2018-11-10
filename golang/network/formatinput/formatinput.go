package main
import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "strings"
)
func sayhelloName(w http.ResponseWriter, r *http.Request) {
    r.ParseForm() //parse url argc
    fmt.Println(r.Form) //log of server
    fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
    fmt.Fprintf(w, "Hello astaxie!") //response to client
}

func login(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method) //get method post or get
    if r.Method == "GET" {
        t, _ := template.ParseFiles("login.gtpl")
        t.Execute(w, nil)
    } else {
        fmt.Println("username:", r.Form["username"])
        fmt.Println("password:", r.Form["password"])
    }
}

func main() {
    http.HandleFunc("/", sayhelloName) 
    http.HandleFunc("/login", login) 
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
