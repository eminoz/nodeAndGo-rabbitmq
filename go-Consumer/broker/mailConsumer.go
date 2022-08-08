package broker

import (
	"encoding/json"
	"fmt"
	"github.com/eminoz/rabbitmq/model"
	"log"
)

func mailConsumer() {
	ch, conn := RabbitConnection()

	q, err := ch.QueueDeclare(
		"mailqueue", // name
		true,        // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)

	if err != nil {
		log.Fatalf("%s: %s", "Failed to declare a queue", err)
	}
	defer conn.Close()
	defer ch.Close()

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	if err != nil {
		log.Fatalf("%s: %s", "Failed to register consumer", err)
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			user := model.User{}
			json.Unmarshal(d.Body, &user)
			fmt.Println(user.Email, "'e mail yollandı ")
			d.Ack(false)
		}
	}()
	<-forever
}