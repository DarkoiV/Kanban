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
    var newList taskList;
    newList.fromJSON(rq.Body)

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

// (DELETE) Delete list
func (bh handler) DeleteList(rw http.ResponseWriter, rq *http.Request) {
    rqVars := mux.Vars(rq)
    boardID := rqVars["boardID"]
    listID  := rqVars["listID"]

    result := bh.db.Where("id = ?", listID).Where("board_id = ?", boardID).Delete(&taskList{})
    if result.RowsAffected == 0 {
        bh.l.Println("List with ID:", listID, "on board", boardID,  "does not exist")
        rw.WriteHeader(http.StatusNotFound)
        return
    }
}
