package mq

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

func DeclareQueue(name QueueName) (*amqp.Connection, *amqp.Channel, error) {
	conn, err := connect()
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a Channel")

	_, err = ch.QueueDeclare(
		string(name),
		false,
		false,
		false,
		false,
		nil)

	return conn, ch, err
}

func DeclareExchange(exname ExchangeName) (*amqp.Connection, *amqp.Channel, error) {
	conn, err := connect()
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a Channel")

	err = ch.ExchangeDeclare(
		string(exname), // name
		"direct", // type
		true, // durable
		false, // auto-deleted
		false, // internal
		false, // no-wait
		nil, // arguments
	)

	return conn, ch, err
}

func connect() (*amqp.Connection, error) {
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

	return amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", username, password, host, port))
}

func failOnError(err error, msg string)  {
	if err != nil {
		log.Fatalf("%s : %s", msg, err)
	}
}