package rabbitmq

import (
	mq "gin-api/config/rabbitmq"
	"log"
)

func init() {
	Client := mq.Client
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

	msgs, err := Client.Consume(
		Declare.Name, // queue
		"",           // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	failOnError(err, "Failed to register a consumer")

	_ = make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s \n ", d.Body)
		}
	}()

}
