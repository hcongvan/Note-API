package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hcongvan/Note-API/control"
	"github.com/hcongvan/Note-API/model"
)

func InitApp() {
	app := control.InitRoute()
	app.Use(gin.Logger())
	model.InitDB("test.db")
	model.InitRabbitMQ()
	app.Run(":8080")
}
