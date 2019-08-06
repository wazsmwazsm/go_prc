// nats-streaming server 通过 clientId + channel + durableName 标示一个持久订阅（参考源码 server.go）
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
	// ConnectWait 控制获取 tcp 连接的超时时间
	sc, err := stan.Connect(clusterID, clientID,
		stan.NatsURL("nats://172.16.129.5:4222"), stan.ConnectWait(time.Second))
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	options := []stan.SubscriptionOption{
		stan.DurableName("my-durable"),
		stan.AckWait(20 * time.Second),
	}

	ch := make(chan stan.Msg)
	// Durable subscribe
	sub, err := sc.Subscribe("foo", func(m *stan.Msg) {
		ch <- *m
	}, options...)
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Close()

	for {
		m := <-ch
		fmt.Printf("%s %v %v %s\n", m.Subject, m.Sequence, m.Timestamp, m.Data)
	}
}
