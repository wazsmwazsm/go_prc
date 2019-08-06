package main

import (
	"log"

	stan "github.com/nats-io/stan.go"
)

const clusterID = "eventbus"
const clientID = "myIdpub"

func main() {
	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL("nats://172.16.129.5:4222"))
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	sc.Publish("foo", []byte("Hello World!"))
}
