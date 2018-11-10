package main 

import (
    "fmt"
    "net/http"
    "strconv"
)

var i int

func myfunc(w http.ResponseWriter, r *http.Request) {
    i++
    fmt.Fprintf(w, strconv.Itoa(i))
}

func main() {
    http.HandleFunc("/",myfunc)
    http.ListenAndServe(":8080",nil)
}
