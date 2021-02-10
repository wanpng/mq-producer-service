package mq

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

// DeclareQueue declare queue
func DeclareQueue(name QueueName) (*amqp.Connection, *amqp.Channel, error) {
	conn, err := Connect()
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

// DeclareExchange declare exchange
func DeclareExchange(exname ExchangeName) (*amqp.Connection, *amqp.Channel, error) {
	conn, err := Connect()
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a Channel")

	err = ch.ExchangeDeclare(
		string(JobseekerEx), // name
		"direct",            // type
		true,                // durable
		false,               // auto-deleted
		false,               // internal
		false,               // no-wait
		nil,                 // arguments
	)

	err = ch.ExchangeDeclare(
		string(JobsEx), // name
		"direct",       // type
		true,           // durable
		false,          // auto-deleted
		false,          // internal
		false,          // no-wait
		nil,            // arguments
	)

	return conn, ch, err
}

// Connect connect to rabbitmq server
func Connect() (*amqp.Connection, error) {
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

	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", username, password, host, port))

	time.Sleep(500 * time.Microsecond)

	return conn, err
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s : %s", msg, err)
	}
}
