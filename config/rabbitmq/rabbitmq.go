package rabbitmq

import (
	"gin-api/config/yaml"
	"github.com/streadway/amqp"
	"log"
)

var Client *amqp.Channel

func init() {
	config := yaml.Conf().RabbitMq
	url := "amqp://" + config.Username + ":" + config.Password + "@" + config.Host + ":" + config.Port + "/"
	//"amqp://guest:guest@localhost:5672/"
	conn, err := amqp.Dial(url)
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
