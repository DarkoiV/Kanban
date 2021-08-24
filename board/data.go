package board

import (
	"encoding/json"
	"io"
    "time"
)

// Task data
type task struct {
    ID              uint        `json:"id"`
    Pos             uint        `json:"pos"`
    Description     string      `json:"description"`
    CreatedAt       time.Time   `json:"createdAt"`
    ListID          uint        `json:"-"`
}

// List of tasks
type taskList struct {
    ID              uint        `json:"id"`
    Pos             uint        `json:"pos"`
    Title           string      `json:"title"`
    Tasks           []task      `json:"tasks"   gorm:"foreignKey:ListID"`
    BoardID         uint        `json:"-"`
}

// Board data 
type board struct {
    ID              uint        `json:"id"`
    Name            string      `json:"name"`
    Lists           []taskList  `json:"lists"   gorm:"foreignKey:BoardID"`
}

// Create JSON representation from board, and write it to io.Writer
func (b *board) toJSON(w io.Writer) error {
    e := json.NewEncoder(w)
    return e.Encode(b)
}

// Create board struct from JSON
func (b *board) fromJSON(r io.Reader) error {
    e := json.NewDecoder(r)
    return e.Decode(b)
}

// Create JSON representation from task list, and write it to io.Writer
func (l *taskList) toJSON(w io.Writer) error {
    e := json.NewEncoder(w)
    return e.Encode(l)
}

// Create JSON representation from task, and write it to io.Writer
func (t *task) toJSON(w io.Writer) error {
    e := json.NewEncoder(w)
    return e.Encode(t)
}
