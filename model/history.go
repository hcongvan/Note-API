package model

import (
	"encoding/json"
	"time"

	"github.com/streadway/amqp"
)

type HistoryLog struct {
	ID       uint        `json:"id"`
	Previous interface{} `json:"previous"`
	Current  interface{} `json:"current"`
	Status   int         `json:"status"`
}

func PublishMessage(body HistoryLog) error {
	bodyByte, errp := json.Marshal(body)
	if errp != nil {
		return errp
	}
	err := channel.Publish(
		"note",
		"note.vanhc",
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
