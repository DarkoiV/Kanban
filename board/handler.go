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

// GET board
func (lh *boardHandler) GetBoard(rw http.ResponseWriter, rq *http.Request) {
    // Log action
    lh.l.Println("GET board from:", rq.URL)

    // Write content type
    rw.Header().Set("Content-Type", "application/json")

    // TODO Load data and convert to JSON
    fmt.Fprint(rw, "{\"msg\":\"TEST\" }")

}

// GET lists
func (lh *boardHandler) GetLists(rw http.ResponseWriter, rq *http.Request) {
    // Log action
    lh.l.Println("GET lists of tasks from:", rq.URL)

    // Load data and convert to JSON 
    taskList := getMockData()
    err := taskList.toJSON(rw)

    // Check for errors during JSON convertion
    if err != nil {
        lh.l.Println("Error", err)
        http.Error(rw, "Error with getting list", http.StatusInternalServerError)
        return
    }

    rw.Header().Set("Content-Type", "application/json")
}
