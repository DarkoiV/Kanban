package data

// Struct holding task data
type Task struct {
    ID              int         `json:"id"`
    Description     string      `json:"description"`
    DueDate         string      `json:"dueDate"`
}

// List of tasks
type TaskList struct {
    ID              int         `json:"id"`
    Title           string      `json:"title"`
    Tasks           []Task      `json:"tasks"`
}
