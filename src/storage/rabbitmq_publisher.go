package storage

import (
	"encoding/json"
        "log"


        amqp "github.com/rabbitmq/amqp091-go"
)

type Message struct {
	Type    string
	Payload interface{}
}

func failOnError(err error, msg string) {
        if err != nil {
                log.Panicf("%s: %s", msg, err)
        }
}

// Produce adds a message into the rabbitMQ queue
func Produce(m Message) {
        conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
        failOnError(err, "Failed to connect to RabbitMQ")
        defer conn.Close()

        ch, err := conn.Channel()
        failOnError(err, "Failed to open a channel")
        defer ch.Close()

	q, err := ch.QueueDeclare(
		"crud",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	b := &m.Payload
	body, err := json.Marshal(b)
	failOnError(err, "Failed to read message body")

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Type:        m.Type,
			Body:        body,
		},
	)
	failOnError(err, "Failed to publish a message")
}
