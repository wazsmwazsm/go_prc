package main

import (
	stan "github.com/nats-io/stan.go"
	"log"
	"strings"
)

const ClusterID = "eventbus"
const ClientID = "TestEventPub"
const ChannelName = "test"

var Servers = []string{
	"nats://172.16.129.5:4222",
	"nats://172.16.129.6:4222",
	"nats://172.16.129.7:4222",
}

func main() {
	// connect nats streaming cluster
	sc, err := stan.Connect(ClusterID, ClientID, stan.NatsURL(strings.Join(Servers, ",")))
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	sc.Publish(ChannelName, []byte("Hello World!"))
}
