package control

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hcongvan/Note-API/model"
)

type RequestFormat struct {
	User  string
	Title string
	Note  string
}

func GetAllNote(c *gin.Context) {

}

func CreateNote(c *gin.Context) {
	var newNote model.NoteModel
	if err := c.BindJSON(&newNote); err != nil {
		fmt.Println("chet cmn roi ae oi")
	}
	newNote.CreateNote(model.DB)
	c.JSON(http.StatusOK, gin.H{"result": newNote})
}

func UpdateNote(c *gin.Context) {

}

func DeleteNote(c *gin.Context) {

}
