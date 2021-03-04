package mq

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

type MessageBody struct {
	Data []byte
	Type string
}

type Message struct {
	Queue         string
	ReplyTo       string
	ContentType   string
	CorrelationID string
	Priority      uint8
	Body          MessageBody
}

type Connection struct {
	name     string
	conn     *amqp.Connection
	channel  *amqp.Channel
	exchange string
	queues   []string
	err      chan error
}

var (
	connectionPool = make(map[string]*Connection)
)

func NewConnection(name string, exchange string, queues ...string) *Connection {
	if c, ok := connectionPool[name]; ok {
		return c
	}

	c := &Connection{
		exchange: exchange,
		queues:   queues,
		err:      make(chan error),
	}

	connectionPool[name] = c
	return c
}

func GetConnection(name string) *Connection {
	return connectionPool[name]
}

func (c *Connection) Connect() error {
	var err error
	var username, password, host string

	if username = viper.GetString("mqusername"); username == "" {
		username = os.Getenv("mqusername")
	}
	if password = viper.GetString("mqpassword"); password == "" {
		password = os.Getenv("mqpassword")
	}
	if host = viper.GetString("mqhost"); host == "" {
		host = os.Getenv("mqhost")
	}
	port := viper.GetInt("mqport")

	c.conn, err = amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", username, password, host, port))

	if err != nil {
		return fmt.Errorf("Error in creating RabbitMQ connection with: %s", err.Error())
	}

	go func() {
		<-c.conn.NotifyClose(make(chan *amqp.Error))
		c.err <- errors.New("Connection closed")
	}()

	c.channel, err = c.conn.Channel()

	if err != nil {
		return fmt.Errorf("Error in channel: %s", err)
	}

	if err := c.channel.ExchangeDeclare(
		c.exchange,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return fmt.Errorf("Error in Exchange declare: %s", err)
	}
	return nil
}

func (c *Connection) BindQueue() error {
	for _, q := range c.queues {
		queueName := fmt.Sprintf("%s:%s", q, os.Getenv("ENV"))
		if _, err := c.channel.QueueDeclare(queueName, true, false, false, false, nil); err != nil {
			return fmt.Errorf("Error in declaring queue %s: %s", q, err)
		}

		if err := c.channel.QueueBind(queueName, queueName, c.exchange, false, nil); err != nil {
			return fmt.Errorf("Error in binding queue %s: %s", q, err)
		}
	}

	return nil
}

func (c *Connection) Reconnect() error {
	if err := c.Connect(); err != nil {
		return err
	}
	if err := c.BindQueue(); err != nil {
		return err
	}

	return nil
}
