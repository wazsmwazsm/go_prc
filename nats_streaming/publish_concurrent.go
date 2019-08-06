package main

import (
	"log"
	stan "github.com/nats-io/stan.go"
	"time"
	"strconv"
)

const clusterID = "eventbus"
const clientID = "myIdpub"

func main() {
	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL("nats://192.168.5.103:4222"))
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	for i := 0; i < 100; i++ {
		go sc.Publish("foo", []byte("Hello World! " + strconv.Itoa(i)))
	}
	
	time.Sleep(1e9 * 5)
}
