package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats"
)

type Company struct {
	Name      string
	Valuation int
}

type Paddle struct {
	Name         string
	Power        int
	Manufacturer *Company
}

func Publish(ec *nats.EncodedConn) gin.HandlerFunc {
	return func(c *gin.Context) {
		company := &Company{
			Name:      "Double Happiness",
			Valuation: 10000000,
		}
		paddle := &Paddle{
			Name:         "Hurricane Long",
			Power:        9000,
			Manufacturer: company,
		}
		ec.Publish("paddle", paddle)
	}
}

func main() {
	nc, _ := nats.Connect(nats.DefaultURL)
	c, _ := nats.NewEncodedConn(nc, nats.GOB_ENCODER)

	c.Subscribe("paddle", func(subject, reply string, p *Paddle) {
		log.Println("Subject: ", subject)
		log.Println("Reply: ", reply)
		log.Println("Paddle Name: ", p.Name)
		log.Println("Paddle Power: ", p.Power)
		log.Println("Paddle Manufacturer: ", p.Manufacturer.Name)
		log.Println("Paddle Manufacturer Valuation: ", p.Manufacturer.Valuation)
	})

	r := gin.Default()
	r.GET("/", Publish(c))
	r.Run(":3000")
}
