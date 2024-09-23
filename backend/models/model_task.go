package models

// Task - Object representing a Task
type Task struct {

	// indicates if a taks is completed or not
	Completed bool `json:"completed,omitempty"`

	// description of the task
	Description string `json:"description"`

	// id of the taks
	Id string `json:"id,omitempty"`
}
