package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("go run rabbit receivers")

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

	messages, err := channelRabbitMQ.Consume(
		"go.queue.uuid", // queue name
		"",              // consumer
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		panic(err)
	}

	for message := range messages {
		fmt.Printf("Received message: %s\n", message.Body)
	}
}
