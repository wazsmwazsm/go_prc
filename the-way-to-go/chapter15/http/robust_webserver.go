package main

import (
	"io"
	"log"
	"net/http"
)

const form = `<html><body><form action="#" method="post" name="bar">
		<input type="text" name="in"/>
		<input type="submit" value="Submit"/>
	</form></html></body>`

type HandleFunc func(http.ResponseWriter, *http.Request)

func SimpleServer(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>hello, world</h1>")
}

func FormServer(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	switch request.Method {
	case "GET":
		io.WriteString(w, form)
	case "POST":
		in := request.FormValue("in")
		if in == "" {
			io.WriteString(w, "in is empty")
			panic("in is empty")
		}
		io.WriteString(w, in)
	}
}

// 这个相当于做了一个 panic 捕捉的中间件
func logPanic(f HandleFunc) HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("[%v] caught panic: %v", r.RemoteAddr, err)
			}
		}()

		f(w, r)
	}
}

func main() {
	http.HandleFunc("/test1", logPanic(SimpleServer))
	http.HandleFunc("/test2", logPanic(FormServer))

	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
