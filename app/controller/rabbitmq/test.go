package rabbitmq

import (
	mq "gin-api/config/rabbitmq"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

func Test1(c *gin.Context) {
	client := mq.Client
	q, err := client.QueueDeclare(
		"test1", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	mq.FailOnError(err, "Failed to declare a queue")

	body := "锄禾日当午!"
	err = client.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	mq.FailOnError(err, "Failed to publish a message")
}
