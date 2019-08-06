// nats-streaming server 通过 clientId + channel + durableName 标示一个持久订阅（参考源码 server.go）
package main

import (
	"fmt"
	stan "github.com/nats-io/stan.go"
	"log"
)

const clusterID = "eventbus"
const clientID = "myIdsub"

func main() {
	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL("nats://172.16.129.5:4222"))
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	ch := make(chan stan.Msg)
	// Durable subscribe
	// nats-streaming server 通过 clientId + channel + durableName 标示一个持久订阅（参考源码 server.go）
	// 下面的例子里，一个连接 (clientId) 中的同 channel 的两个订阅者，用不同
	// 的 durableName 来标示两个不通的持久化订阅
	// 这也是 durableName 的意义
	// 每个订阅者取消订阅后不会影响其它订阅者
	sub, err := sc.Subscribe("foo", func(m *stan.Msg) {
		ch <- *m
	}, stan.DurableName("myIdsub1"))
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Close()

	sub1, err := sc.Subscribe("foo", func(m *stan.Msg) {
		ch <- *m
	}, stan.DurableName("myIdsub2"))
	if err != nil {
		log.Fatal(err)
	}
	defer sub1.Close()

	for {
		m := <-ch
		fmt.Printf("%s %v %v %s\n", m.Subject, m.Sequence, m.Timestamp, m.Data)
	}
}
