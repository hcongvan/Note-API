package model

import (
	"encoding/json"
	"os"
	"time"

	"github.com/streadway/amqp"
)

type HistoryLog struct {
	ID       uint        `json:"note_id"`
	Previous interface{} `json:"previous"`
	Current  interface{} `json:"current"`
	Type     int         `json:"type"`
}

func PublishMessage(body HistoryLog) error {
	bodyByte, errp := json.Marshal(body)
	if errp != nil {
		return errp
	}
	err := channel.Publish(
		os.Getenv("RABBIT_EXCHANGE"),
		os.Getenv("RABBIT_ROUTINGKEY"),
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Timestamp:   time.Now(),
			Body:        bodyByte,
		},
	)
	return err
}
