package main

import (
	"fmt"
	"net/http"
)

type hello struct {
	routes map[string]map[string]func(w http.ResponseWriter, r *http.Request)
}

func NewHello() *hello {
	routes := make(map[string]map[string]func(w http.ResponseWriter, r *http.Request))
	return &hello{routes}
}

// 实现了 Handler 接口 ServeHTTP 里面来处理路由注册、匹配等
func (h *hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	path := r.URL.Path
	handler, ok := h.routes[path][method]
	if !ok {
		// w.WriteHeader(http.StatusNotFound)
		// fmt.Fprint(w, "404 Not Found!")
		http.NotFound(w, r) // 简介的方式
		return
	}
	handler(w, r)
}

func (h *hello) request(path, method string, handler func(w http.ResponseWriter, r *http.Request)) {
	methodMap := make(map[string]func(w http.ResponseWriter, r *http.Request))
	methodMap[method] = handler
	h.routes[path] = methodMap
}
func (h *hello) Get(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	h.request(path, "GET", handler)
}

func (h *hello) Post(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	h.request(path, "POST", handler)
}

func main() {
	h := NewHello()
	h.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello root!")
	})
	h.Get("/aa", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "aa!")
	})
	h.Post("/post", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "post!")
	})

	http.ListenAndServe(":9999", h)
}
