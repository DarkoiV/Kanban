package board

import (
	"encoding/json"
	"io"
	"time"
)

// Task ///////////////////////////////////////////////////////////////////////
type task struct {
	ID          uint      `json:"id"`
	Pos         uint      `json:"pos"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	ListID      uint      `json:"-"`
	BoardID     uint      `json:"-"`
}

func (t *task) toJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(t)
}

func (t *task) fromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(t)
}

// List of tasks //////////////////////////////////////////////////////////////
type taskList struct {
	ID      uint   `json:"id"`
	Pos     uint   `json:"pos"`
	Title   string `json:"title"`
	Tasks   []task `json:"tasks"   gorm:"foreignKey:ListID;constraint:OnDelete:CASCADE"`
	BoardID uint   `json:"-"`
}

func (l *taskList) toJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(l)
}

func (l *taskList) fromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(l)
}

// Board //////////////////////////////////////////////////////////////////////
type board struct {
	ID    uint       `json:"id"`
	Name  string     `json:"name"`
	Lists []taskList `json:"lists"   gorm:"foreignKey:BoardID;constraint:OnDelete:CASCADE"`
	Tasks []task     `json:"-"       gorm:"foreignKey:BoardID;constraint:OnDelete:CASCADE"`
}

func (b *board) toJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(b)
}

func (b *board) fromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(b)
}
