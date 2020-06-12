package models

import "reflect"

// ToDoStatus type
type ToDoStatus int

// TodoStatus enum definitions
const (
	Open ToDoStatus = iota
	InProgress
	Closed
)

// Return ToDoStatus string
func (s ToDoStatus) String() string {
	switch s {
	case 0:
		return "Open"
	case 1:
		return "In Progress"
	case 2:
		return "Closed"
	default:
		return "Unknown"
	}
}

// ToDo struct
type ToDo struct {
	ID          int64      `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	CreatedAt   string     `json:"createdAt"`
	UpdatedAt   string     `json:"updatedAt"`
	Status      ToDoStatus `json:"status"`
	Tags        []Tag      `json:"tags,omitempty"`
}

// Tag struct
type Tag struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
}

// IsEmpty will return if todo is empty
func (x ToDo) IsEmpty() bool {
	return reflect.DeepEqual(x, ToDo{})
}

// IsEmpty will return if tag is empty
func (x Tag) IsEmpty() bool {
	return reflect.DeepEqual(x, Tag{})
}
