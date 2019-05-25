package main

import (
	"fmt"
	"net/http"
)

// 自定义路由
type MyMux struct {

}
// 实现 Handler 接口
func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		// 访问 / 时运行 sayHelloName handler
		sayHelloName(w, r)
		return
	}
	// 其他路由则运行 not found handler
	http.NotFound(w, r)
	return
}


func main() {
	mux := &MyMux{}
	http.ListenAndServe(":9090", mux) // 使用外部路由
}	


func sayHelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello My Route!")
}
