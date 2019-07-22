package main

import (
	stan "github.com/nats-io/stan.go"
	"log"
)

func main() {
	sc, err := stan.Connect("test-cluster", "myId1")
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	sc.Publish("foo", []byte("Hello World aaa"))
}
