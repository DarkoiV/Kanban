package handlers

import (
	"log"
	"net/http"

	"github.com/darkoiv/Kanban-back-go/data"
)

// Board handler
type boardHandler struct {
    l *log.Logger
}

// Create task list handler with logger
func NewBoardHandler (l *log.Logger) *boardHandler {
    return &boardHandler{l}
}

// GET lists
func (lh *boardHandler) GetLists(rw http.ResponseWriter, rq *http.Request) {
    // Log action
    lh.l.Println("GET lists of tasks from:", rq.URL)

    // Write content type
    rw.Header().Set("Content-Type", "application/json")

    // Load data and convert to JSON 
    taskList := data.GetMockData()
    err := taskList.ToJSON(rw)

    // Check for errors during JSON convertion
    if err != nil {
        lh.l.Println("Error", err)
        http.Error(rw, "Error with getting list", http.StatusInternalServerError)
        return
    }

}
