package control

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hcongvan/Note-API/model"
)

func GetAllNote(c *gin.Context) {
	// var notes []model.NoteModel
	notes := model.NoteModel{}.GetAllNote()
	c.JSON(http.StatusOK, ResponseMsg{Code: OK, Result: notes})
}

func CreateNote(c *gin.Context) {
	var newNote model.NoteModel
	if err := c.BindJSON(&newNote); err != nil {
		c.JSON(http.StatusBadRequest, ResponseMsg{Code: Error, Result: Message{Msg: DataWrong}})
		return
	}
	newNote, err := newNote.CreateNote()
	if err {
		c.JSON(http.StatusBadRequest, ResponseMsg{Code: Error, Result: Message{Msg: BadData}})
		return
	}
	c.JSON(http.StatusOK, ResponseMsg{Code: OK, Result: newNote})
}

func UpdateNote(c *gin.Context) {
	var updateNote model.NoteModel
	if err := c.BindJSON(&updateNote); err != nil {
		c.JSON(http.StatusBadRequest, ResponseMsg{Code: Error, Result: Message{Msg: DataWrong}})
		return
	}
	id, errp := strconv.Atoi(c.Param("id"))
	if errp != nil {
		c.JSON(http.StatusBadRequest, ResponseMsg{Code: Error, Result: Message{Msg: IdWrong}})
		return
	}
	updateNote.ID = uint(id)
	updateNote, err := updateNote.UpdateNotebyId()
	if err {
		c.JSON(http.StatusBadRequest, ResponseMsg{Code: Error, Result: Message{Msg: DontExist}})
		return
	}
	c.JSON(http.StatusOK, ResponseMsg{Code: OK, Result: updateNote})
}

func DeleteNote(c *gin.Context) {
	var deleteNote model.NoteModel
	id, errp := strconv.Atoi(c.Param("id"))
	if errp != nil {
		c.JSON(http.StatusBadRequest, ResponseMsg{Code: Error, Result: Message{Msg: IdWrong}})
		return
	}
	deleteNote.ID = uint(id)
	err := deleteNote.DeleteNotebyId()
	if err {
		c.JSON(http.StatusBadRequest, ResponseMsg{Code: Error, Result: Message{Msg: DontExist}})
		return
	}
	c.JSON(http.StatusOK, ResponseMsg{Code: OK, Result: Message{Msg: NoteDelete}})
}
