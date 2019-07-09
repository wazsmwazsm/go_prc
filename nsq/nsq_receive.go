package main

import (
	"fmt"
	"time"
	"github.com/nsqio/go-nsq"
)

type ConsumerT struct{}

func (*ConsumerT) HandleMessage(msg *nsq.Message) error {
	fmt.Println(string(msg.Body))
	return nil
}

func init() {
	topic, channel, address := "test", "test_chan2", "127.0.0.1:4161"

	cfg := nsq.NewConfig()
	cfg.LookupdPollInterval = time.Second 
	con, err := nsq.NewConsumer(topic, channel, cfg)
	if err != nil {
		panic(err)
	}

	con.SetLogger(nil, 0) // 关掉日志
	con.AddHandler(&ConsumerT{})

	if err := con.ConnectToNSQLookupd(address); err != nil {
		panic(err)
	}
}

func main() {
	// 主 goroutine 常驻, 
	for {
		time.Sleep(time.Second * 10)
	}
}

