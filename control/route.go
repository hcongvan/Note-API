package control

import (
	"github.com/gin-gonic/gin"
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
)

func InitRoute() *gin.Engine {
	_route := gin.Default()

	_route.GET("/notes", GetAllNote)
	_route.POST("/notes", CreateNote)
	_route.PUT("/notes/:id", UpdateNote)
	_route.DELETE("notes/:id", DeleteNote)

	return _route
}
