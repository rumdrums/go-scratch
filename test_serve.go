package main

import (
	"net"
	"net/http"
    //"github.com/gorilla/mux"
)

func main() {
    listener, _ := net.Listen("tcp", "0.0.0.0:8080")
    fs := http.FileServer(http.Dir("static"))
	//router := mux.NewRouter()
    //router.Handle("/", fs)
	http.Handle("/", fs)
	http.Serve(listener, nil)
}
