package main

import (
	//"io"
	"net/http"
)

/*
net/http/server.go:
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
*/

type helloHandler struct{} //helloHandler implement interface Handler

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func main() {
	http.Handle("/", &helloHandler{})
	http.ListenAndServe(":12345", nil)
}
