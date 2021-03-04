package mq

import (
	"fmt"

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

	if err := c.channel.Publish(c.exchange, m.Queue, false, false, p); err != nil {
		return fmt.Errorf("Error in publishing: %s", err)
	}

	return nil
}
