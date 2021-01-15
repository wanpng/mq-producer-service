package mq

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

func DeclareQueue(name QueueName) (*amqp.Connection, *amqp.Channel, error) {
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
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a Channel")

	_, err = ch.QueueDeclare(
		fmt.Sprintf("%s-%s", string(name), os.Getenv("ENV")),
		false,
		false,
		false,
		false,
		nil)

	return conn, ch, err
}

func failOnError(err error, msg string)  {
	if err != nil {
		log.Fatalf("%s : %s", msg, err)
	}
}