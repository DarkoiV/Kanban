package board

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// (POST) Create list
func (bh handler) CreateList(rw http.ResponseWriter, rq *http.Request) {
	rqVars := mux.Vars(rq)
	boardID := rqVars["boardID"]

	var newList taskList
	if err := newList.fromJSON(rq.Body); err != nil {
		bh.l.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	boardIDNumber, err := strconv.Atoi(boardID)
	if err != nil {
		bh.l.Print("Error getting board ID from URL", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	newList.BoardID = uint(boardIDNumber)
	bh.db.Create(&newList)

	// Send response
	if err := newList.toJSON(rw); err != nil {
		bh.l.Println("Error with JSON marshal,", err)
		rw.WriteHeader(http.StatusInternalServerError)
	}
}

// (PATCH) Update lists
func (bh handler) UpdateList(rw http.ResponseWriter, rq *http.Request) {
	rqVars := mux.Vars(rq)
	boardID := rqVars["boardID"]
	listID := rqVars["listID"]

	var reqList taskList
	if err := reqList.fromJSON(rq.Body); err != nil {
		bh.l.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	listIDNum, err := strconv.Atoi(listID)
	if err != nil {
		bh.l.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	result := bh.db.Where("id = ?", listID).Where("board_id = ?", boardID).Find(&taskList{})
	if result.Error != nil || result.RowsAffected == 0 {
		bh.l.Println("List with ID:", listID, "on board", boardID, "does not exist")
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	result = bh.db.Updates(&reqList)
	if result.Error != nil || result.RowsAffected == 0 {
		bh.l.Println(result.Error)
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	tasks := reqList.Tasks
	for _, utask := range tasks {
		utask.ListID = uint(listIDNum)
		utaskPos := strconv.Itoa(int(utask.Pos))
		result := bh.db.Model(&task{}).Where("id = ?", utask.ID).Updates(map[string]interface{}{"pos": utaskPos, "list_id": utask.ListID})
		if result.Error != nil || result.RowsAffected == 0 {
			bh.l.Println(result.Error)
		}
	}

	if err := reqList.toJSON(rw); err != nil {
		bh.l.Println("Error with JSON marshal,", err)
		rw.WriteHeader(http.StatusInternalServerError)
	}
}

// (DELETE) Delete list
func (bh handler) DeleteList(rw http.ResponseWriter, rq *http.Request) {
	rqVars := mux.Vars(rq)
	boardID := rqVars["boardID"]
	listID := rqVars["listID"]

	result := bh.db.Where("id = ?", listID).Where("board_id = ?", boardID).Delete(&taskList{})
	if result.RowsAffected == 0 {
		bh.l.Println("List with ID:", listID, "on board", boardID, "does not exist")
		rw.WriteHeader(http.StatusNotFound)
		return
	}
}
