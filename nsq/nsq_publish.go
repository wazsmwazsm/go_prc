package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
	"time"
	"encoding/json"
)

// nsqd node
var ips = []string{
	"http://127.0.0.1:4153/pub?topic=test",
	"http://127.0.0.1:4155/pub?topic=test",
}

type Msg struct {
	Timestamp int64
	Msg       string
}

func main() {
	// 生成时间戳，用来给消费者判断重复消费
	s, _ := json.Marshal(Msg{time.Now().UnixNano(), "Hello 啊!"})
	ch := make(chan string)
	for _, ip := range ips {
		go send(ip, string(s), ch)
	}

	for i := 0; i < len(ips); i++ {
		fmt.Println(<- ch)
	}
}

func send(ip string, msg string, ch chan string) {
	
	resp, err := http.Post(ip, "application/json", strings.NewReader(msg))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	ch <- ip + ":" + string(body)
}
