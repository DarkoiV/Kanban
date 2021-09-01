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

    var reqTask task
    if err := reqTask.fromJSON(rq.Body); err != nil {
        bh.l.Print("Error with JSON unmarshall,", err)
        rw.WriteHeader(http.StatusInternalServerError)
        return
    }

    reqTask.ListID = func(listID string) uint{
        ID, _ := strconv.Atoi(listID)
        return uint(ID)
    }(listID)

    result := bh.db.Where("id = ?", listID).Where("board_id = ?", boardID).Find(&taskList{})
    if result.Error != nil || result.RowsAffected == 0 {    
        bh.l.Println("List with ID:", listID, "on board", boardID,  "does not exist")
        rw.WriteHeader(http.StatusNotFound)
        return
    }

    result = bh.db.Create(&reqTask)
    if result.Error != nil || result.RowsAffected == 0 {
        bh.l.Println("Error when creating new task", result.Error)
        rw.WriteHeader(http.StatusNotFound)
        return
    }

    reqTask.toJSON(rw)
}

// (PUT) update task
func (bh handler) UpdateTask(rw http.ResponseWriter, rq *http.Request) {
    rqVars := mux.Vars(rq)
    boardID := rqVars["boardID"]
    listID := rqVars["listID"]
    taskID := rqVars["taskID"]

    var reqTask task
    if err := reqTask.fromJSON(rq.Body); err != nil {
        bh.l.Print("Error with JSON unmarshall,", err)
        rw.WriteHeader(http.StatusInternalServerError)
        return
    }

    // Make sure it is on correct endpoint
    if reqID, err := strconv.Atoi(taskID); uint(reqID) != reqTask.ID || err != nil {
        bh.l.Print(reqTask.ID, err)
        bh.l.Print("Wrong endpoint for this task")
        rw.WriteHeader(http.StatusInternalServerError)
        return
    }

    // Check if list is on board
    result := bh.db.Where("id = ?", listID).Where("board_id = ?", boardID).Find(&taskList{})
    if result.Error != nil || result.RowsAffected == 0 {    
        bh.l.Println("List with ID:", listID, "on board", boardID,  "does not exist")
        rw.WriteHeader(http.StatusNotFound)
        return
    }

    // Check if task is on list, and update if so
    result = bh.db.Where("id = ?", taskID).Where("list_id = ?", listID).Updates(reqTask)
    if result.Error != nil || result.RowsAffected == 0 {
        bh.l.Println("Task with ID:", taskID, "on list", listID,  "does not exist")
        rw.WriteHeader(http.StatusNotFound)
        return
    }

    reqTask.toJSON(rw)
}

// (DELETE) task
func (bh handler) DeleteTask(rw http.ResponseWriter, rq *http.Request) {

}
