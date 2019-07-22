package main

import (
	nats "github.com/nats-io/nats.go"
	"log"
	"time"
	"fmt"
)

func main() {
	// connect
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Subscribe
	sub, err := nc.SubscribeSync("foo")
	if err != nil {
		log.Fatal(err)
	}

	// Wait for a message
	msg, err := sub.NextMsg(10 * time.Second)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Reply: %s", msg.Data)
}
