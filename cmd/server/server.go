package main

import (
	"log"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal("failed to connect to rabitmq")
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("failed to open a channel")
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   //  delete when unused
		false,   //  exlusive
		false,   // no-wait
		nil,     //argument
	)
	if err != nil {
		log.Fatal("failed to declare a queue")
	}

	msg := "Test Message"
	if err := ch.Publish(
		"",     // exhange
		q.Name, // routing key
		false,  // mendatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		}); err != nil {
		log.Fatal("failed to declare a queue")
	}
}
