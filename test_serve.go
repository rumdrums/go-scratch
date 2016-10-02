package main

import (
	"net"
	"net/http"
)

func main() {
    listener, _ := net.Listen("tcp", "0.0.0.0:8080")
	http.Serve(listener, nil)
}
