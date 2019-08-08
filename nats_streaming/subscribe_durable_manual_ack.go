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
		stan.SetManualAckMode(),
	}

	ch := make(chan stan.Msg)
	// Durable subscribe
	sub, err := sc.Subscribe("foo", func(m *stan.Msg) {
		fmt.Printf("响应前可以拿到消息, 但是 server 还不知道已经交付: %s %v %v %s\n", m.Subject, m.Sequence, m.Timestamp, m.Data)
		// 手动 ack, 更安全, 可以保证在业务完成后 ack, 出错则关闭连接, 下次再 redelivery，
		// 这样可以在处理一些逻辑后再 ack，保证这次消息和消费是原子操作。
		time.Sleep(time.Second * 3)
		m.Ack()
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
