package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

// http 升级需要 Upgrade Connect 之类的 header 信息, 使用 websocket 包封装好的 Upgrader 来实现
var upgrader = websocket.Upgrader{}

func main() {
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/", home)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "home.html")
}

func echo(w http.ResponseWriter, r *http.Request) {
	// 获取升级协议后的 ws 连接
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer ws.Close()

	for { // loop 来接受信息
		mt, msg, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
			continue
		}

		if err := ws.WriteMessage(mt, append([]byte("hello "), msg...)); err != nil {
			log.Println(err)
		}
	}
}
