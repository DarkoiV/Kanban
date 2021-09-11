package board

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// (POST) create board
func (bh handler) CreateBoard(rw http.ResponseWriter, rq *http.Request) {
	var newBoard board
	newBoard.fromJSON(rq.Body)
	bh.db.Create(&newBoard)

	// Send response
	if err := newBoard.toJSON(rw); err != nil {
		bh.l.Println("Error with JSON marshal, ", err)
		rw.WriteHeader(http.StatusInternalServerError)
	}

}

// (GET) list of boards
func (bh handler) GetBoards(rw http.ResponseWriter, rq *http.Request) {
	listItems := []struct {
		ID   uint    `json:"id"`
		Name string  `json:"name"`
	}{}

	result := bh.db.Model(&board{}).Find(&listItems)
	if result.Error != nil {
		bh.l.Println(result.Error)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send response
	e := json.NewEncoder(rw)
	if err := e.Encode(listItems); err != nil {
		bh.l.Println("Error with JSON marshal, ", err)
		rw.WriteHeader(http.StatusInternalServerError)
	}

}

// (GET) board
func (bh handler) GetBoard(rw http.ResponseWriter, rq *http.Request) {
	rqVars := mux.Vars(rq)
	boardID := rqVars["boardID"]

	var resBoard board

	// Load board from database
	result := bh.db.Where("id = ?", boardID).Preload("Lists").Preload("Lists.Tasks").First(&resBoard)
	if result.Error != nil {
		bh.l.Println(result.Error)
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	// Send response
	if err := resBoard.toJSON(rw); err != nil {
		bh.l.Println("Error with JSON marshal, ", err)
		rw.WriteHeader(http.StatusInternalServerError)
	}
}

// (DELETE) board
func (bh handler) DeleteBoard(rw http.ResponseWriter, rq *http.Request) {
	rqVars := mux.Vars(rq)
	boardID := rqVars["boardID"]

	result := bh.db.Where("id = ?", boardID).Delete(&board{})
	if result.RowsAffected == 0 || result.Error != nil {
		bh.l.Println("Board with ID:", boardID, "does not exist", " or: ", result.Error)
		rw.WriteHeader(http.StatusNotFound)
		return
	}
}
