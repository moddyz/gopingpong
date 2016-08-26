package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats"
)

func Ping(nc *nats.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Ping")
		nc.Publish("ping", []byte("Ping!"))
	}
}

func main() {
	nc, _ := nats.Connect(nats.DefaultURL)

	r := gin.Default()
	r.GET("/", Ping(nc))
	r.Run(":3000")
}
