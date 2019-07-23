package main

import (
	nats "github.com/nats-io/nats.go"
	"log"
	"fmt"
	"sync"
)

func main() {
	// connect
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	wg := sync.WaitGroup{}
	wg.Add(1)

	// Subscribe
	if _, err := nc.Subscribe("foo", func (m *nats.Msg) {
		wg.Done()
		fmt.Printf("%s", m.Data)
	}); err != nil {
		log.Fatal(err)
	}

	wg.Wait()
}
