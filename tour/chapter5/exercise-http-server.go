package main

import (
	"fmt"
	"log"
	"net/http"
)

type String string
// 实现 Handler 接口
func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}
// 实现 Handler 接口
func (s *Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s%s %s", s.Greeting, s.Punct, s.Who)
}

func main() {
	// 设置路由
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
	err := http.ListenAndServe("0.0.0.0:4000", nil)

	if err != nil {
		log.Fatal(err)
	}
}
