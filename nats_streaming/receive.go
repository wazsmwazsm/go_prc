package main

import (
	stan "github.com/nats-io/stan.go"
	"log"
	"fmt"
	"os"
	"os/signal"
)

func PrintMsg(m *stan.Msg) {
	fmt.Printf("%v %t %s\n", m.Sequence, m.Redelivered, m.Data)
}

func main() {
	sc, err := stan.Connect("test-cluster", "myId1")
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	startOpt := stan.DeliverAllAvailable()

	sub, err := sc.QueueSubscribe("foo", "fooquene", PrintMsg, startOpt)
	if err != nil {
		sc.Close()
		log.Fatal(err)
	}

	sub.Unsubscribe()

	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		for range signalChan {
			sub.Unsubscribe()
			sc.Close()
			cleanupDone <- true
		}
	}()
	<- cleanupDone
}
