package board

import (
	"encoding/json"
	"io"
)

// Card data
type card struct {
    ID              uint        `json:"id"`
    Pos             uint        `json:"pos"`
    Description     string      `json:"description"`
    DueDate         string      `json:"dueDate"`
    ListID          uint        `json:"-"`

}

// List of cards
type cardList struct {
    ID              uint        `json:"id"`
    Pos             uint        `json:"pos"`
    Title           string      `json:"title"`
    Cards           []card      `json:"cards" gorm:"foreignKey:ListID"`
    BoardID         uint        `json:"-"`
}

// Board data 
type board struct {
    ID              uint        `json:"id"`
    Name            string      `json:"name"`
    Lists           []cardList  `json:"lists" gorm:"foreignKey:BoardID"`
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

// Create JSON representation from card list, and write it to io.Writer
func (l *cardList) toJSON(w io.Writer) error {
    e := json.NewEncoder(w)
    return e.Encode(l)
}

// Create JSON representation from card, and write it to io.Writer
func (c *card) toJSON(w io.Writer) error {
    e := json.NewEncoder(w)
    return e.Encode(c)
}
