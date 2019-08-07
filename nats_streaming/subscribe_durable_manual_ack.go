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
	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL("nats://172.16.129.5:4222"))
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	options := []stan.SubscriptionOption{
		stan.DurableName("my-durable"),
		stan.AckWait(20 * time.Second),
		stan.SetManualAckMode(),
	}

	ch := make(chan stan.Msg)
	// Durable subscribe
	sub, err := sc.Subscribe("foo", func(m *stan.Msg) {
		m.Ack() // 手动 ack, 更安全, 可以保证在业务完成后 ack, 出错则关闭连接, 下次 redelivery
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
