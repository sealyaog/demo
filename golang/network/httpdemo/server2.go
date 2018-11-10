package main

import (
	"io"
	"net/http"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func main() {
	http.HandleFunc("/", helloHandler) //yaog: HandleFunc is a adptor, we do need to implement a struct
	http.ListenAndServe(":12345", nil)
}
