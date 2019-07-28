package natsstreaming

import (
	"log"
	stan "github.com/nats-io/stan.go"
	"httpserve/config"
)

var Conn stan.Conn

func init() {
	clusterId, _ := config.Get("clusterID").(string)
	natsURL, _ := config.Get("natsStreamingUrl").(string)
	sc, err := stan.Connect(clusterId, "testpub", stan.NatsURL(natsURL))
	Conn = sc
	if err != nil {
		log.Fatal(err)
	}
}
