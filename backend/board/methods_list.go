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
		writeMessageJSON(rw, "Server error")
		return
	}

	boardIDNumber, err := strconv.Atoi(boardID)
	if err != nil {
		bh.l.Print("Error getting board ID from URL", err)
		rw.WriteHeader(http.StatusInternalServerError)
		writeMessageJSON(rw, "Server error")
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

	// Create struct containing data for updates
	var reqList taskList
	if err := reqList.fromJSON(rq.Body); err != nil {
		bh.l.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
		writeMessageJSON(rw, "Server error")
		return
	}

	// get list ID from URL
	listIDNum, err := strconv.Atoi(listID)
	if err != nil {
		bh.l.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
		writeMessageJSON(rw, "Server error")
		return
	}

	// Update list
	result := bh.db.Model(&taskList{}).
		Where("id = ?", listID).
		Where("board_id = ?", boardID).
		Updates(map[string]interface{}{"title": reqList.Title, "pos": reqList.Pos})
	if result.Error != nil || result.RowsAffected == 0 {
		bh.l.Println(result.Error, "or if nil list not found")
		rw.WriteHeader(http.StatusNotFound)
		writeMessageJSON(rw, "Not found")
		return
	}

	// Update pos and associations with list of tasks
	tasks := reqList.Tasks
	completed := make(chan bool, len(tasks))
	for _, taskForUpdate := range tasks {
		go func(taskForUpdate task) {
			taskForUpdate.ListID = uint(listIDNum)

			// Querry
			result := bh.db.Model(&task{}).
				Where("id = ?", taskForUpdate.ID).
				Where("board_id = ?", boardID).
				Updates(map[string]interface{}{"pos": taskForUpdate.Pos, "list_id": taskForUpdate.ListID})

			if result.Error != nil {
				bh.l.Println(result.Error)
			}

			completed <- true
		}(taskForUpdate)
	}

	// Wait for completion of task updates
	for range tasks {
		<-completed
	}

	// Send request as response if all ok
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
		writeMessageJSON(rw, "Not found")
		return
	}

	writeMessageJSON(rw, "Deleted succesfully")
}
