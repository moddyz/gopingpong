package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats"
)

func main() {
	nc, _ := nats.Connect(nats.DefaultURL)
	nc.Subscribe("ping", func(m *nats.Msg) {
		log.Printf("Got a %s\n", string(m.Data))
		log.Println("Imma Pong!")
	})

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		log.Println("I need a ping")
	})
	r.Run(":3001")
}
