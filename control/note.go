package control

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hcongvan/Note-API/model"
)

func GetAllNote(c *gin.Context) {
	// var notes []model.NoteModel
	limit, errparse := strconv.Atoi(c.Query("limit"))
	if errparse != nil {
		c.JSON(http.StatusOK, ResponseMsg{Code: OK, Result: Message{Msg: errparse.Error()}})
		return
	}
	offset, errparse := strconv.Atoi(c.Query("offset"))
	if errparse != nil {
		c.JSON(http.StatusOK, ResponseMsg{Code: OK, Result: Message{Msg: errparse.Error()}})
		return
	}
	notes, err := model.NoteModel{}.GetAllNote(offset, limit)
	if err != nil {
		c.JSON(http.StatusOK, ResponseMsg{Code: OK, Result: Message{Msg: BadConn}})
		return
	}
	c.JSON(http.StatusOK, ResponseMsg{Code: OK, Result: notes})
}

func CreateNote(c *gin.Context) {
	var newNote model.NoteModel
	if err := c.BindJSON(&newNote); err != nil {
		c.JSON(http.StatusBadRequest, ResponseMsg{Code: Error, Result: Message{Msg: DataWrong}})
		return
	}
	newNote, err := newNote.CreateNote()
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseMsg{Code: Error, Result: Message{Msg: BadData}})
		return
	}
	c.JSON(http.StatusOK, ResponseMsg{Code: OK, Result: newNote})
}

func UpdateNote(c *gin.Context) {
	var updateNote model.NoteModel
	var upNote map[string]interface{}
	if err := c.BindJSON(&upNote); err != nil {
		c.JSON(http.StatusBadRequest, ResponseMsg{Code: Error, Result: Message{Msg: DataWrong}})
		return
	}
	id, errp := strconv.Atoi(c.Param("id"))
	if errp != nil {
		c.JSON(http.StatusBadRequest, ResponseMsg{Code: Error, Result: Message{Msg: IdWrong}})
		return
	}
	updateNote.Id = uint(id)
	updateNote, err := updateNote.UpdateNote(upNote)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseMsg{Code: Error, Result: Message{Msg: err.Error()}})
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
	deleteNote.Id = uint(id)
	err := deleteNote.DeleteNote()
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseMsg{Code: Error, Result: Message{Msg: DontExist}})
		return
	}
	c.JSON(http.StatusOK, ResponseMsg{Code: OK, Result: Message{Msg: NoteDelete}})
}
