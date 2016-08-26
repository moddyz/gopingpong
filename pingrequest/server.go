package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats"
)

func PingRequest(nc *nats.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Pinging...")
		msg, err := nc.Request("ping", []byte("Ping!"), 5*time.Second)
		if err != nil {
			log.Println("Error on ping request, ", err)
			c.AbortWithError(http.StatusRequestTimeout, err)
			return
		}
		log.Println("Received a ", string(msg.Data))
		c.String(http.StatusOK, "Received a %s", msg.Data)
	}
}

func main() {
	nc, _ := nats.Connect(nats.DefaultURL)

	r := gin.Default()
	r.GET("/", PingRequest(nc))
	r.Run(":3000")
}
