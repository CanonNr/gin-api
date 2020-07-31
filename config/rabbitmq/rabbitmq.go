package rabbitmq

import (
	"github.com/streadway/amqp"
	"log"
)

var Client *amqp.Channel

func init() {
	conn, err := amqp.Dial("amqp://admin:123457@172.16.2.30:5672/")
	FailOnError(err, "Failed to connect to RabbitMQ")
	//defer conn.Close()
	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	//defer ch.Close()
	Client = ch
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
