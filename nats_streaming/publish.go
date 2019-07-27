package main

import (
	"log"

	stan "github.com/nats-io/stan.go"
)

const clusterID = "eventbus"
const clientID = "myIdpub"

func main() {
	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL("nats://192.168.5.103:6222"))
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	sc.Publish("foo", []byte("Hello World!"))
}
