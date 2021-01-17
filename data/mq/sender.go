package mq

import (
	"encoding/json"
	"log"
	"fmt"
	"os"
	
	"github.com/streadway/amqp"
)

func SendMQ(model interface{}, ch *amqp.Channel, q QueueName) {
	b, err := json.MarshalIndent(&model, "", "\t")

	if err != nil {
		log.Fatalf("Cannot marshal json data")
	}

	err = ch.Publish(
		"",
		fmt.Sprintf("%s-%s", string(q), os.Getenv("ENV")),
		false,
		false,
		amqp.Publishing {
			DeliveryMode: amqp.Persistent,
			ContentType: "application/json",
			Body: []byte(b),
		})
}


func SendMQEx(model interface{}, ch *amqp.Channel, e ExchangeName, r RoutingKey) {
	b, err := json.MarshalIndent(&model, "", "\t")

	if err != nil {
		log.Fatalf("Cannot marshal json data")
	}

	err = ch.Publish(
		string(e),
		string(r),
		false,
		false,
		amqp.Publishing {
			DeliveryMode: amqp.Persistent,
			ContentType: "application/json",
			Body: []byte(b),
		})
}