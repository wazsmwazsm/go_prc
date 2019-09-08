package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello,"+req.URL.Path[len("/hello/"):])
}

func ShortHelloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "shorthello,"+strings.ToUpper(req.URL.Path[len("/shorthello/"):]))
}

func main() {
	http.HandleFunc("/hello/", HelloHandler)
	http.HandleFunc("/shorthello/", ShortHelloHandler)
	err := http.ListenAndServe("localhost:9999", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
