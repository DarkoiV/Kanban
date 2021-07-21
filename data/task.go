package data

// Struct holding task data
type Task struct {
    ID              int         `json:"id"`
    Description     string      `json:"description"`
    DueDate         string      `json:"dueDate"`
}
