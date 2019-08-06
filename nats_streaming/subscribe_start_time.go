package main

import (
	"fmt"
	stan "github.com/nats-io/stan.go"
	"log"
	"time"
)

const clusterID = "eventbus"
const clientID = "myIdsub"

func main() {
	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL("nats://192.168.5.101:4222"))
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	ch := make(chan stan.Msg)
	sub, err := sc.Subscribe("foo", func(m *stan.Msg) {
		ch <- *m
	}, stan.StartAtTime(time.Unix(1564234753, 162929938)), stan.AckWait(20*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Close()
	// defer sub.Unsubscribe() // when not durable mode, as same as close
	for {
		m := <-ch
		fmt.Printf("%v %v %s\n", m.Sequence, m.Timestamp, m.Data)
	}
}
