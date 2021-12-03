package control

import (
	"consumer/model"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/streadway/amqp"
)

func ReadMessage(msgs <-chan amqp.Delivery) {
	for d := range msgs {
		log.Printf(" [x] %s", d.Body)
		if (d.ContentType == "application/json") && (d.RoutingKey == os.Getenv("RABBIT_ROUTINGKEY")) {
			var _tmp map[string]interface{}
			err := json.Unmarshal(d.Body, &_tmp)
			if err != nil {
				fmt.Println(err)
			}
			ty := int(_tmp["type"].(float64))
			noteId := uint(_tmp["note_id"].(float64))
			switch ty {
			case 1:
				current := _tmp["current"].(map[string]interface{})
				createHistory(noteId, current, d.Timestamp)
			case 2:
				current := _tmp["current"].(map[string]interface{})
				previous := _tmp["previous"].(map[string]interface{})
				updateHistory(noteId, current, previous, d.Timestamp)
			case 3:
				deleteHistory(noteId, d.Timestamp)
			}
			d.Ack(true)
		}
	}
}

func createHistory(note_id uint, current map[string]interface{}, timeStamp time.Time) {
	var logCur model.NoteHistory
	logCur.CreateAt = timeStamp
	logCur.NoteId = note_id
	logCur.PreviousId = 0
	logCur.User = current["user"].(string)
	logCur.Title = current["title"].(string)
	logCur.Note = current["note"].(string)
	logCur.CreateAt = timeStamp
	logCur.Status = true
	logCur, err := logCur.CreateNoteHistory()
	if err != nil {
		fmt.Println(err)
	}
}

func updateHistory(note_id uint, current map[string]interface{}, previous map[string]interface{}, timeStamp time.Time) {
	var logCur model.NoteHistory
	var logPre model.NoteHistory
	logCur.CreateAt = timeStamp
	logCur.NoteId = note_id

	logCur.User = current["user"].(string)
	logCur.Title = current["title"].(string)
	logCur.Note = current["note"].(string)
	logCur.CreateAt = timeStamp
	logCur.Status = true
	logPre.NoteId = note_id
	logPre.User = current["user"].(string)
	logPre.Title = current["title"].(string)
	logPre.Note = current["note"].(string)
	logPre.Status = true
	logPre, err := logPre.GetNoteHistory()
	if err != nil {
		fmt.Println(err)
	}
	logCur.PreviousId = logPre.Id
	logCur, errCreate := logCur.CreateNoteHistory()
	if errCreate != nil {
		fmt.Println(errCreate)
	}
}

func deleteHistory(note_id uint, timeStamp time.Time) {
	var logCur model.NoteHistory
	logCur.NoteId = note_id
	err := logCur.DeleteNoteHistory()
	if err != nil {
		fmt.Println(err)
	}
}
