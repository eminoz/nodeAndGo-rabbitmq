package broker

import (
	"github.com/streadway/amqp"
	"log"
)

func RabbitConnection() (*amqp.Channel, *amqp.Connection) {
	conn, err := amqp.Dial("amqp://" + "eminoz" + ":" + "eminoz" + "@" + "localhost" + ":" + "5672" + "/")

	if err != nil {
		log.Fatalf("%s: %s", "Failed to connect to RabbitMQ", err)
	}

	ch, err := conn.Channel()

	if err != nil {
		log.Fatalf("%s: %s", "Failed to open a channel", err)
	}
	return ch, conn
}
