package board

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// Board handler
type boardHandler struct {
    l *log.Logger
    db *gorm.DB
}

// Create board handler
func NewHandler (l *log.Logger, db *gorm.DB) *boardHandler {
    l.Println("Auto migrating database")
    db.AutoMigrate(&board{}, &cardList{}, &card{})
    return &boardHandler{l, db}
}

///// BOARD METHODS ///////////////////////////////////////////////////////////////////////////////

// (GET) board
func (bh *boardHandler) GetBoard(rw http.ResponseWriter, rq *http.Request) {
    bh.l.Println("GET board from:", rq.URL)
    rw.Header().Set("Content-Type", "application/json")

    rqVars := mux.Vars(rq)
    boardID := rqVars["boardID"]

    var resBoard board;

    // Check if exists 
    result := bh.db.Limit(1).Where("id = ?", boardID).Find(&resBoard)
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
func (bh *boardHandler) CreateBoard(rw http.ResponseWriter, rq *http.Request) {
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

///// LISTS METHODS ///////////////////////////////////////////////////////////////////////////////

// (GET) lists on board
func (bh *boardHandler) GetLists(rw http.ResponseWriter, rq *http.Request) {


    rw.Header().Set("Content-Type", "application/json")
}

// (PATCH) Update order of cards in lists
func (bh *boardHandler) UpdateOrder(rw http.ResponseWriter, rq *http.Request) {

    rw.Header().Set("Content-Type", "application/json")
}

// (DELETE) list from board
func (bh *boardHandler) DeleteList(rw http.ResponseWriter, rq *http.Request) {

    rw.Header().Set("Content-Type", "application/json")
}

///// CARDS METHODS ///////////////////////////////////////////////////////////////////////////////

// (POST) Create new card
func (bh *boardHandler) PostCard(rw http.ResponseWriter, rq *http.Request) {

    rw.Header().Set("Content-Type", "application/json")
}

// (PUT) update card
func (bh *boardHandler) UpdateCard(rw http.ResponseWriter, rq *http.Request) {

    rw.Header().Set("Content-Type", "application/json")
}

// (DELETE) card
func (bh *boardHandler) DeleteCard(rw http.ResponseWriter, rq *http.Request) {

    rw.Header().Set("Content-Type", "application/json")
}
