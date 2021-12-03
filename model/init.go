package model

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var channel *amqp.Channel

func InitDB(db_name string) {
	_db, err := gorm.Open(sqlite.Open(db_name), &gorm.Config{})
	if err != nil {
		fmt.Printf("failed to connect database")
	}
	_db.AutoMigrate(&NoteModel{})
	db = _db
}

func InitRabbitMQ() {
	conn, err := amqp.Dial("amqp://rabbit-admin:xHc2zUkq4PZLeQ2C@42.119.139.251:31678/")
	if err != nil {
		log.Fatalf("%s: %s", err, "fail to connect")
	}
	ch, errch := conn.Channel()
	if errch != nil {
		log.Fatalf("%s: %s", errch, "fail when open channel")
	}
	channel = ch
}
