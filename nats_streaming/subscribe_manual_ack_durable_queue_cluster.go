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

var clientID = flag.String("clientid", "", "client id")
var sleep = flag.Int64("sleep", 0, "sleep time")

func main() {
	flag.Parse()
	if *clientID == "" {
		log.Fatal("client id is empty")
	}
	// connect nats streaming cluster
	sc, err := stan.Connect(ClusterID, *clientID, stan.NatsURL(strings.Join(Servers, ",")))
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	ch := make(chan *stan.Msg)
	msgHandler := func(m *stan.Msg) {
		// ack 前可以拿到消息，消费方可以确定消费逻辑成功后再 ack
		// 加入一个延迟, 开启手动 ack, 模拟 queue 模式中不同消费者可能 ack
		// 次序不同导致 queue 中的消息被消费无法保证时间次序的场景
		time.Sleep(time.Duration(*sleep) * time.Second)
		m.Ack()
		ch <- m
	}
	// queue Durable subscribe
	// queue 的模式下，由于 server 会 pending 未 ack 的 subscriber 连接等待重新交付
	// 所以没有办法保证整个 queue 的消费数据是顺序的
	sub, err := sc.QueueSubscribe(ChannelName, QueueName, msgHandler,
		stan.DurableName(DurableName),
		stan.SetManualAckMode(),
	)
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
			// 捕捉到信号后，执行 defer，显示关闭 subscriber 或 sc 连接（否则 server 会认为 subscriber 连接没有关闭
			// ，发送还会发到该 subscriber 连接上，只是没有 ack，会 pending 在 server，直到 server 判定失活
			// 关闭该 subscriber 连接，pending 的数据才会重新交付给其它 subscriber 上）
			return
		}
	}
}
