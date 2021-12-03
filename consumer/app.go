package main

import (
	"consumer/control"
	"consumer/model"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	model.InitDB("test.db")
	control.InitRabbitMQ()

}
