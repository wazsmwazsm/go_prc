package main

import (
	"flag"
	"fmt"
	stan "github.com/nats-io/stan.go"
	"log"
	"os"
	"os/signal"
	"strings"
	"time"
)

const ClusterID = "eventbus"

// queue durable 模式下，server 通过 queueName + durableName 来标识持久化订阅
const QueueName = "testQueue"
const DurableName = "test"
const ChannelName = "test"

var Servers = []string{
	"nats://172.16.129.5:4222",
	"nats://172.16.129.6:4222",
	"nats://172.16.129.7:4222",
}

var ClientID = flag.String("clientid", "", "client id")

func main() {
	flag.Parse()
	if *ClientID == "" {
		log.Fatal("client id is empty")
	}
	// connect nats streaming cluster
	sc, err := stan.Connect(ClusterID, *ClientID, stan.NatsURL(strings.Join(Servers, ",")))
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	ch := make(chan *stan.Msg)
	msgHandler := func(m *stan.Msg) {
		ch <- m
	}
	// Durable subscribe
	sub, err := sc.QueueSubscribe(ChannelName, QueueName, msgHandler,
		stan.DurableName(DurableName), stan.MaxInflight(1), stan.AckWait(time.Second))
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Close()

	c := make(chan os.Signal, 1)
	signal.Notify(c)

	for {
		select {
		case m := <-ch:
			fmt.Printf("channel: %s, seq: %v, time: %v, msg: %s\n", m.Subject, m.Sequence, m.Timestamp, m.Data)
		case s := <-c:
			fmt.Println(s)
			return
		}
	}
}
