package control

import (
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

var channel *amqp.Channel

func InitRabbitMQ() {
	connectionString := fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		os.Getenv("RABBIT_USER"),
		os.Getenv("RABBIT_PASSWD"),
		os.Getenv("RABBIT_HOST"),
		os.Getenv("RABBIT_PORT"),
	)
	conn, err := amqp.Dial(connectionString)
	if err != nil {
		log.Fatalf("%s: %s", err, "fail to connect")
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("%s: %s", err, "fail to open channel")
	}
	channel = ch

	msgs, err := ch.Consume(
		os.Getenv("RABBIT_QUEUE"),
		os.Getenv("RABBIT_CONSUMERNAME"),
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("%s: %s", err, "fail to register a consumer")
	}
	done := make(chan bool)
	go ReadMessage(msgs)
	<-done
}
