package handlers

import (
	"log"
	"net/http"

	"github.com/darkoiv/Kanban-back-go/data"
)

// Task list handler
type ListsHandler struct {
    l *log.Logger
}

// Create task list handler with logger
func NewListsHandler (l *log.Logger) *ListsHandler {
    return &ListsHandler{l}
}

// GET lists
func (lh *ListsHandler) GetLists(rw http.ResponseWriter, rq *http.Request) {
    // Load data and convert to JSON 
    taskList := data.GetMockData()
    lh.l.Println("GET lists of tasks from:", rq.URL);
    err := taskList.ToJSON(rw)

    // Check for errors during JSON convertion
    if err != nil {
        lh.l.Println("Error", err)
        http.Error(rw, "Error with getting list", http.StatusInternalServerError)
        return
    }
}
