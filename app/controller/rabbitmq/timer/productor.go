package timer

import (
	"gin-api/app/controller/rabbitmq"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func Set(ctx *gin.Context) {
	var err error
	client := rabbitmq.Client

	body := "test"
	// 将消息发送到延时队列上
	err = client.Publish(
		"",           // exchange 这里为空则不选择 exchange
		"test_delay", // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
			Expiration:  "5000", // 设置五秒的过期时间
		})
	failOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s", body)
}
