package board

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// (POST) Create new task
func (bh handler) PostTask(rw http.ResponseWriter, rq *http.Request) {
	rqVars := mux.Vars(rq)
	boardID := rqVars["boardID"]
	listID := rqVars["listID"]

	// Get task from request
	var reqTask task
	if err := reqTask.fromJSON(rq.Body); err != nil {
		bh.l.Print("Error with JSON unmarshall,", err)
		rw.WriteHeader(http.StatusInternalServerError)
		writeMessageJSON(rw, "Server error")
		return
	}

	reqTask.ListID = func(listID string) uint {
		ID, _ := strconv.Atoi(listID)
		return uint(ID)
	}(listID)

	// Check for list on board
	result := bh.db.Where("id = ?", listID).Where("board_id = ?", boardID).Find(&taskList{})
	if result.Error != nil || result.RowsAffected == 0 {
		bh.l.Println("List with ID:", listID, "on board", boardID, "does not exist")
		rw.WriteHeader(http.StatusNotFound)
		writeMessageJSON(rw, "Not found")
		return
	}

	reqTask.BoardID = func(boardID string) uint {
		ID, _ := strconv.Atoi(boardID)
		return uint(ID)
	}(boardID)

	// Create task in DB
	result = bh.db.Create(&reqTask)
	if result.Error != nil {
		bh.l.Println("Error when creating new task", result.Error)
		rw.WriteHeader(http.StatusInternalServerError)
		writeMessageJSON(rw, "Server error")
		return
	}

	// Respond with task with ID
	if err := reqTask.toJSON(rw); err != nil {
		bh.l.Println("Error with JSON marshal,", err)
		rw.WriteHeader(http.StatusInternalServerError)
	}
}

// (PUT) update task
func (bh handler) UpdateTask(rw http.ResponseWriter, rq *http.Request) {
	rqVars := mux.Vars(rq)
	boardID := rqVars["boardID"]
	listID := rqVars["listID"]
	taskID := rqVars["taskID"]

	// Get task from request
	var reqTask task
	if err := reqTask.fromJSON(rq.Body); err != nil {
		bh.l.Print("Error with JSON unmarshall,", err)
		rw.WriteHeader(http.StatusInternalServerError)
		writeMessageJSON(rw, "Server error")
		return
	}

	// Make sure it is on correct endpoint
	if reqID, err := strconv.Atoi(taskID); uint(reqID) != reqTask.ID || err != nil {
		bh.l.Print(reqTask.ID, err)
		bh.l.Print("Wrong endpoint for this task")
		rw.WriteHeader(http.StatusInternalServerError)
		writeMessageJSON(rw, "Server error")
		return
	}

	// Check if task is on list and board, and update if so
	result := bh.db.Where("id = ?", taskID).Where("list_id = ?", listID).Where("board_id = ?", boardID).Updates(reqTask)
	if result.Error != nil {
		bh.l.Println("Task not updated:", result.Error)
		rw.WriteHeader(http.StatusNotFound)
		writeMessageJSON(rw, "Not found")
		return
	}

	// Respond with updated task
	if err := reqTask.toJSON(rw); err != nil {
		bh.l.Println("Error with JSON marshal,", err)
		rw.WriteHeader(http.StatusInternalServerError)
	}
}

// (DELETE) task
func (bh handler) DeleteTask(rw http.ResponseWriter, rq *http.Request) {

}
