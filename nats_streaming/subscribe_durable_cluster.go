package main

import (
	"fmt"
	stan "github.com/nats-io/stan.go"
	"log"
	"strings"
)

const ClusterID = "eventbus"
const ClientID = "TestEventSub"

// nats-streaming server 通过 clientId + channel + durableName 标示一个持久订阅（参考源码 server.go）
const DurableName = "test"
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

	ch := make(chan *stan.Msg)
	msgHandler := func(m *stan.Msg) {
		ch <- m
	}
	// Durable subscribe
	sub, err := sc.Subscribe(ChannelName, msgHandler, stan.DurableName(DurableName))
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Close()

	for {
		m := <-ch
		fmt.Printf("channel: %s, seq: %v, time: %v, msg: %s\n", m.Subject, m.Sequence, m.Timestamp, m.Data)
	}
}
