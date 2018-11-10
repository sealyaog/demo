package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

func fakeDial(proto, addr string)(conn net.Conn, err error) {         
    return net.Dial("unix", "/var/run/docker.sock")                               
} 

func main() {

	tr := &http.Transport{Dial: fakeDial,}

	client := &http.Client{Transport: tr}

	resp, err := client.Get("http://v1.24/containers/5e0cf91fb8cc/stats?stream=0")
	//resp, err := client.Get("http://v1.24/version")

    if err != nil {
		log.Fatal(err)
	}


	robots, err := ioutil.ReadAll(resp.Body)

	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", robots)
}
