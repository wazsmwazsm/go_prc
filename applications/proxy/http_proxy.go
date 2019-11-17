package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type reqERR struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type Transport struct{}

func (Transport) RoundTrip(r *http.Request) (*http.Response, error) {

	resp, err := http.DefaultTransport.RoundTrip(r)
	resp.Header.Del("Content-Length") // 要修改 body, 先删除旧的响应中的实体长度
	return resp, err
}

func main() {
	target, _ := url.Parse("http://127.0.0.1:8081/")
	log.Printf("forwarding to -> %s\n", target)

	proxy := httputil.NewSingleHostReverseProxy(target)

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// 响应修改回调
		proxy.ModifyResponse = func(resp *http.Response) error {
			buf, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil
			}

			code := 0
			errMsg := ""

			if resp.StatusCode != 200 {
				reqErr := &reqERR{}
				err := json.Unmarshal(buf, reqErr)
				if err == nil { // is json
					code = reqErr.Code
					errMsg = reqErr.Message

					respStr := fmt.Sprintf("{\"code\":%v,\"error_message\":\"%s\", \"data\":[]}", code, errMsg)
					buf = []byte(respStr)
				}

			} else {
				respStr := fmt.Sprintf("{\"code\":%v,\"error_message\":\"\", \"data\":%s}", code, buf)
				buf = []byte(respStr)
			}

			resp.Body = ioutil.NopCloser(bytes.NewReader(buf))
			return nil
		}
		proxy.Transport = &Transport{} // 从代理服务到后端服务的传输过程

		proxy.ServeHTTP(w, req) // 处理代理
	})

	log.Fatal(http.ListenAndServe(":8082", nil))
}
