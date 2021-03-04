package mq

import (
	"fmt"
	"os"

	"github.com/streadway/amqp"
)

func (c *Connection) Publish(m Message) error {
	select {
	case err := <-c.err:
		if err != nil {
			c.Reconnect()
		}
	default:
	}

	p := amqp.Publishing{
		Headers:       amqp.Table{"type": m.Body.Type},
		ContentType:   m.ContentType,
		CorrelationId: m.CorrelationID,
		Body:          m.Body.Data,
		ReplyTo:       m.ReplyTo,
	}

	queueName := fmt.Sprintf("%s:%s", m.Queue, os.Getenv("ENV"))

	if err := c.channel.Publish(c.exchange, queueName, false, false, p); err != nil {
		return fmt.Errorf("Error in publishing: %s", err)
	}

	return nil
}
