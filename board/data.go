package board

import (
	"encoding/json"
	"io"
)

// Card data
type card struct {
    ID              int         `json:"id"`
    Pos             int         `json:"pos"`
    Description     string      `json:"description"`
    DueDate         string      `json:"dueDate"`
}

// List of cards
type cardList struct {
    ID              int         `json:"id"`
    Pos             int         `json:"pos"`
    Title           string      `json:"title"`
    Cards           []card      `json:"cards"`
}

// Board data 
type board struct {
    ID              int         `json:"id"`
    Name            string      `json:"name"`
    Lists           []cardList  `json:"lists"`
}

// Create JSON representation from board, and write it to io.Writer
func (b *board) toJSON(w io.Writer) error {
    e := json.NewEncoder(w)
    return e.Encode(b)
}
