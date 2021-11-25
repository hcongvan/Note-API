package control

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

type ResponseMsg struct {
	Code   int         `json:"code"`
	Result interface{} `json:"result"`
}

type Message struct {
	Msg string `json:"msg"`
}

const (
	OK    int = 0
	Error     = 1
)
const (
	DataWrong  string = "data body wrong"
	IdWrong           = "ID must be number"
	DontExist         = "Record don't exist"
	NoteDelete        = "Note deleted"
	BadData           = "Bad data"
	BadConn           = "Bad connection"
)

var channel *amqp.Channel
var rabbitmq *amqp.Connection

func InitRoute() *gin.Engine {
	_route := gin.Default()

	_route.GET("/notes", GetAllNote)
	_route.POST("/notes", CreateNote)
	_route.PUT("/notes/:id", UpdateNote)
	_route.DELETE("notes/:id", DeleteNote)

	return _route
}

func InitRabbitMQ() {
	conn, err := amqp.Dial("amqp://rabbit-admin:xHc2zUkq4PZLeQ2C@42.119.139.251:31678/")
	if err != nil {
		log.Fatalf("%s: %s", err, "fail to connect")
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("%s: %s", err, "fail to open channel")
	}
	ch.Close()
	channel = ch
	rabbitmq = conn
}
