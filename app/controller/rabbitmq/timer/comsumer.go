package timer

import (
	"gin-api/app/controller/rabbitmq"
	"log"

	"github.com/streadway/amqp"
)

//func failOnError(err error, msg string) {
//	if err != nil {
//		log.Fatalf("%s: %s", msg, err)
//	}
//}

func init() {
	// 建立链接
	var err error
	client := rabbitmq.Client

	// 声明一个主要使用的 exchange
	err = client.ExchangeDeclare(
		"logs",   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	// 声明一个常规的队列, 其实这个也没必要声明,因为 exchange 会默认绑定一个队列
	q, err := client.QueueDeclare(
		"test_logs", // name
		false,       // durable
		false,       // delete when unused
		true,        // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	failOnError(err, "Failed to declare a queue")

	/**
	 * 注意,这里是重点!!!!!
	 * 声明一个延时队列, ß我们的延时消息就是要发送到这里
	 */
	_, errDelay := client.QueueDeclare(
		"test_delay", // name
		false,        // durable
		false,        // delete when unused
		true,         // exclusive
		false,        // no-wait
		amqp.Table{
			// 当消息过期时把消息发送到 logs 这个 exchange
			"x-dead-letter-exchange": "logs",
		}, // arguments
	)
	failOnError(errDelay, "Failed to declare a delay_queue")

	err = client.QueueBind(
		q.Name, // queue name, 这里指的是 test_logs
		"",     // routing key
		"logs", // exchange
		false,
		nil)
	failOnError(err, "Failed to bind a queue")

	// 这里监听的是 test_logs
	msgs, err := client.Consume(
		q.Name, // queue name, 这里指的是 test_logs
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	//forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
		}
	}()

	//log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	//<-forever
}
