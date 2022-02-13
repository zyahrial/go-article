package message_broker

import (
    // "log"
    // "os"
    "fmt"
    // "github.com/gofiber/fiber/v2"
    // "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/streadway/amqp"
)

func Publish(data []byte, queueName string) string{

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

    q, err := channelRabbitMQ.QueueDeclare(
		queueName, // name
		true,    // durable
		false,   // delete when usused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)

    // Handle any errors if we were unable to create the queue
	if err != nil {
		fmt.Println(err)
	}

	e := channelRabbitMQ.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "json/plain",
			Body:        []byte(data),
		})

    if e != nil{
        fmt.Println(e)
    }

    return "command message success"
}