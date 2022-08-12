package broker

import (
	"encoding/json"
	"fmt"
	"github.com/eminoz/rabbitmq/controller"
	"github.com/eminoz/rabbitmq/model"
	"log"
)

func userConsumer(u controller.UserController) {
	ch, conn := RabbitConnection()
	q, err := ch.QueueDeclare(
		"newuser", // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	mailqueue, err := ch.QueueDeclare(
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
	mailqueueConsume, err := ch.Consume(
		mailqueue.Name, // queue
		"",             // consumer
		false,          // auto-ack
		false,          // exclusive
		false,          // no-local
		false,          // no-wait
		nil,            // args
	)

	if err != nil {
		log.Fatalf("%s: %s", "Failed to register consumer", err)
	}
	forever := make(chan bool)
	go func() {
		for d := range mailqueueConsume {
			user := model.User{}
			json.Unmarshal(d.Body, &user)
			fmt.Println(user.Email, "'e mail yollandı ")
			d.Ack(false)
		}

	}()
	go func() {
		for d := range msgs {
			m := &model.User{}
			err2 := json.Unmarshal(d.Body, m)
			if err2 != nil {
				fmt.Println(err2)
			}
			u.SaveUser(m)
			d.Ack(false)
		}
	}()
	<-forever
}
