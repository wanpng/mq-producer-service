package mq

import (
	"fmt"
	"os"

	"github.com/streadway/amqp"
)

func (c *Connection) Consume() (map[string]<-chan amqp.Delivery, error) {
	m := make(map[string]<-chan amqp.Delivery)
	for _, q := range c.queues {
		queueName := fmt.Sprintf("%s:%s", q, os.Getenv("ENV"))
		deliveries, err := c.channel.Consume(queueName, "", false, false, false, false, nil)

		if err != nil {
			return nil, err
		}
		m[q] = deliveries
	}

	return m, nil
}

func (c *Connection) HandleConsumedDeliveries(q string, delivery <-chan amqp.Delivery, fn func(Connection, string, <-chan amqp.Delivery)) {
	for {
		go fn(*c, q, delivery)

		if err := <-c.err; err != nil {
			c.Reconnect()
			deliveries, err := c.Consume()

			if err != nil {
				panic(err)
			}
			delivery = deliveries[q]
		}
	}
}
