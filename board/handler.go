package board

import (
	"fmt"
	"log"
	"net/http"
)

// Board handler
type boardHandler struct {
    l *log.Logger
}

// Create board handler
func NewHandler (l *log.Logger) *boardHandler {
    return &boardHandler{l}
}

///// BOARD METHODS ///////////////////////////////////////////////////////////////////////////////

// (GET) board
func (bh *boardHandler) GetBoard(rw http.ResponseWriter, rq *http.Request) {
    // Log action
    bh.l.Println("GET board from:", rq.URL)

    // Write content type
    rw.Header().Set("Content-Type", "application/json")

    // TODO Load data and convert to JSON
    fmt.Fprint(rw, "{\"msg\":\"TEST\" }")

}

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
