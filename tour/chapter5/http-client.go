package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	response, err := http.Get("http://www.baidu.com")
	if err != nil {
		// handle error
		fmt.Println("Fatal error ", err.Error())
	}
	// 关闭 response
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
}
