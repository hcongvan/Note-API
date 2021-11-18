package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hcongvan/Note-API/control"
)

func InitRoute() *gin.Engine {
	_route := gin.Default()
	_route.GET("/notes", control.GetAllNote)
	_route.POST("/notes", control.CreateNote)
	_route.PUT("/notes/:id", control.UpdateNote)
	_route.DELETE("notes/:id", control.DeleteNote)

	return _route
}
