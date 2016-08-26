package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats"
)

func main() {
	port := flag.Int("port", 3002, "port number")
	flag.Parse()

	nc, _ := nats.Connect(nats.DefaultURL)
	nc.QueueSubscribe("ping", "queued_pongs", func(m *nats.Msg) {
		log.Printf("Got a %s\n", string(m.Data))
		log.Println("Imma Pong!")
	})

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		log.Println("I need a ping")
	})
	r.Run(fmt.Sprintf(":%v", *port))
}
