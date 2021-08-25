package board

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// Board handler
type handler struct {
    l *log.Logger
    db *gorm.DB
}

// Create board handler
func NewHandler (l *log.Logger, db *gorm.DB) *handler {
    l.Println("Auto migrating database")
    db.AutoMigrate(&board{}, &taskList{}, &task{})
    return &handler{l, db}
}

///// BOARD METHODS ///////////////////////////////////////////////////////////////////////////////

// (GET) board
func (bh handler) GetBoard(rw http.ResponseWriter, rq *http.Request) {
    bh.l.Println("GET board from:", rq.URL)
    rw.Header().Set("Content-Type", "application/json")

    rqVars := mux.Vars(rq)
    boardID := rqVars["boardID"]

    var resBoard board;

    // Load board from database
    result := bh.db.Where("id = ?", boardID).Preload("Lists").Preload("Lists.Tasks").First(&resBoard)
    if result.RowsAffected == 0 {
        bh.l.Println("Board with ID:", boardID, "does not exist")
        rw.WriteHeader(http.StatusBadRequest)
        return
    }

    // Send response
    if err := resBoard.toJSON(rw); err != nil {
        bh.l.Println("Error with JSON marshal, ", err)
        rw.WriteHeader(http.StatusInternalServerError)
    }

}

// (POST) create board
func (bh handler) CreateBoard(rw http.ResponseWriter, rq *http.Request) {
    bh.l.Println("POST new board", rq.URL)
    rw.Header().Set("Content-Type", "application/json")

    // Create new
    var newBoard board;
    newBoard.fromJSON(rq.Body);
    bh.db.Create(&newBoard)

    // Send response
    if err := newBoard.toJSON(rw); err != nil {
        bh.l.Println("Error with JSON marshal, ", err)
        rw.WriteHeader(http.StatusInternalServerError)
    }

}

// (DELETE) board
func (bh handler) DeleteBoard(rw http.ResponseWriter, rq *http.Request) {
    bh.l.Println("DELETE board", rq.URL)

    rqVars := mux.Vars(rq)
    boardID := rqVars["boardID"]

    bh.db.Where("id = ?", boardID).Delete(&board{})
}

///// LISTS METHODS ///////////////////////////////////////////////////////////////////////////////

// (GET) lists on board
func (bh handler) GetLists(rw http.ResponseWriter, rq *http.Request) {


    rw.Header().Set("Content-Type", "application/json")
}

// (PATCH) Update lists
func (bh handler) UpdateLists(rw http.ResponseWriter, rq *http.Request) {

    rw.Header().Set("Content-Type", "application/json")
}

// (DELETE) list from board
func (bh handler) DeleteList(rw http.ResponseWriter, rq *http.Request) {

    rw.Header().Set("Content-Type", "application/json")
}

///// CARDS METHODS ///////////////////////////////////////////////////////////////////////////////

// (POST) Create new task
func (bh handler) PostTask(rw http.ResponseWriter, rq *http.Request) {

    rw.Header().Set("Content-Type", "application/json")
}

// (PUT) update task
func (bh handler) UpdateTask(rw http.ResponseWriter, rq *http.Request) {

    rw.Header().Set("Content-Type", "application/json")
}

// (DELETE) task
func (bh handler) DeleteTask(rw http.ResponseWriter, rq *http.Request) {

    rw.Header().Set("Content-Type", "application/json")
}
