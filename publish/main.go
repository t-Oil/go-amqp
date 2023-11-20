package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("go run rabbit publish")

	connectRabbitMQ, err := amqp.Dial("AMQP.CONFIG")
	if err != nil {
		panic(err)
	}
	defer connectRabbitMQ.Close()

	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}
	defer channelRabbitMQ.Close()

	msg := uuid.New().String()

	message := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(msg),
	}

	if err := channelRabbitMQ.Publish(
		"go.queue.generate.uuid", // exchange
		"",                       // queue name
		false,
		false,
		message, // message to publish
	); err != nil {
		panic(err)
	}

	fmt.Print(msg)
}
