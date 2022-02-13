package message_broker

import (
    // "log"
    // "os"
    "fmt"
    // "github.com/gofiber/fiber/v2"
    // "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/kamva/mgm/v3"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"

    "github.com/streadway/amqp"
	models "query/article/models"
)

func Consume(c *gin.Context){

	amqpServerURL := "amqp://rmq:rmq@rabbitmq:5672/"

    connectRabbitMQ, err := amqp.Dial(amqpServerURL)
    if err != nil {
        panic(err)
    }

    channelRabbitMQ, err := connectRabbitMQ.Channel()
    if err != nil {
        panic(err)
    }

	messages, err := channelRabbitMQ.Consume(
		"QueueAddArticle", // queue
		"",       // consumer
		true,     // auto-ack
		false,    // exclusive
		false,    // no-local
		false,    // no-wait
		nil,      // args
	)

	if err != nil {
		fmt.Println(err)
	}

    fmt.Println("Successfully connected to RabbitMQ")
    fmt.Println("Waiting for messages")

    forever := make(chan bool)

    go func() {
        for message := range messages {
            // fmt.Printf(" > Received message: %s\n", message.Body)
			res := message.Body
			var m models.ShowArticle

			if err := json.Unmarshal(res, &m); err != nil {
				fmt.Println(err)
			}
		
			id := m.Id
			author := m.Author
			tittle := m.Tittle
			body := m.Body

			save := models.NewArticle(id, author, tittle, body)
			succ := mgm.Coll(save).Create(save)
            fmt.Println("success add data to query DB:", succ)
			defer connectRabbitMQ.Close()
			defer channelRabbitMQ.Close()
		}
    }()
    <-forever

	c.JSON(http.StatusOK, gin.H{
		"status": "message success consumed",
	})
}