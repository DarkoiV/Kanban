package data

import (
	"encoding/json"
	"io"
	"time"
)

// Struct holding task data
type Task struct {
    ID              int         `json:"id"`
    Pos             int         `json:"pos"`
    Description     string      `json:"description"`
    DueDate         string      `json:"dueDate"`
}

// List of tasks
type TaskList struct {
    ID              int         `json:"id"`
    Pos             int         `json:"pos"`
    Title           string      `json:"title"`
    Tasks           []*Task     `json:"tasks"`
}

// Alias for multiple task lists
type Lists []TaskList

// Create JSON list from struct
func (lists *Lists) ToJSON(w io.Writer) error {
    e := json.NewEncoder(w)
    return e.Encode(lists)
}

// MOCK DATA ////////////////////////////////////////////////////////////

// Get mock data for now
func GetMockData() Lists {
    return mockData;
}

// TMP MOCK DATA
var mockTasks = []*Task{
    {
        ID: 0,
        Description: "This is my first task",
        DueDate: time.Now().Local().String(),
    },
    {
        ID: 1,
        Description: "This is my SECOND task",
        DueDate: time.Now().Local().String(),
    },
    {
        ID: 2,
        Description: "This is my THIRD task",
        DueDate: time.Now().Local().String(),
    },
    {
        ID: 3,
        Description: "I need to implement garbage\n collection ;')'",
        DueDate: time.Now().Local().String(),
    },
}
var mockTasks2 = []*Task{
    {
        ID: 4,
        Description: "Lorem ipsum",
        DueDate: time.Now().Local().String(),
    },
    {
        ID: 5,
        Description: "Take out that trash \n NOW!",
        DueDate: time.Now().Local().String(),
    },
    {
        ID: 6,
        Description: "Read a book",
        DueDate: time.Now().Local().String(),
    },
    {
        ID: 7,
        Description: "We know",
        DueDate: time.Now().Local().String(),
    },
}
var mockData = Lists {
    TaskList{ID: 0, Title: "To Do", Tasks: mockTasks},
    TaskList{ID: 0, Title: "In progress", Tasks: mockTasks2},
}
