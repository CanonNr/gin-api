package rabbitmq

import (
	mq "gin-api/config/rabbitmq"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"log"
)

var (
	Client  *amqp.Channel
	Declare amqp.Queue
)

func init() {
	Client = mq.Client
	q, err := Client.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	mq.FailOnError(err, "Failed to declare a queue")
	Declare = q
}

func Put(c *gin.Context) {
	body := "消息来了.."
	err := Client.Publish(
		"",           // exchange
		Declare.Name, // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	mq.FailOnError(err, "Failed to publish a message")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
