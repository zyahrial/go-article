package message_broker

import (
    // "log"
    // "os"
    "fmt"
    // "github.com/gofiber/fiber/v2"
    // "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/streadway/amqp"
)

func Consume(){

	amqpServerURL := "amqp://rmq:rmq@rabbitmq:5672/"

    connectRabbitMQ, err := amqp.Dial(amqpServerURL)
    if err != nil {
        panic(err)
    }
    defer connectRabbitMQ.Close()

    channelRabbitMQ, err := connectRabbitMQ.Channel()
    if err != nil {
        panic(err)
    }
    defer channelRabbitMQ.Close()

	msgChan, err := channelRabbitMQ.Consume(
		"QueueAddArticle", // queue
		"",       // consumer
		true,     // auto-ack
		false,    // exclusive
		false,    // no-local
		false,    // no-wait
		nil,      // args
	)
	if err != nil {
		return nil, fmt.Errorf("CONSUMING '%s' CHANNEL: %s", chanName, err.Error())
	}

	fmt.Println(string(msgChan))
	return msgChan, nil
}