package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ToDoStatus type
type ToDoStatus int

// TodoStatus enum definitions
const (
	Open ToDoStatus = iota + 1
	InProgress
	Closed
)

// Return ToDoStatus string
func (s ToDoStatus) String() string {
	switch s {
	case 1:
		return "Open"
	case 2:
		return "In Progress"
	case 3:
		return "Closed"
	default:
		return "Unknown"
	}
}

// ToDo struct
type ToDo struct {
	ID          string
	Title       string
	Description string
	createdAt   int
	updatedAt   int
	state       ToDoStatus
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		results := []ToDo{
			ToDo{"abc123", "first todo", "this is the first item", 1591559402263, 1591559402263, Open},
			ToDo{"def456", "second todo", "this is the second item", 1591559412263, 1591559412263, Open},
			ToDo{"ghi789", "third todo", "this is the third item", 1591559422263, 1591559422263, Open},
		}
		encoder := json.NewEncoder(w)
		if err := encoder.Encode(results); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/open", func(w http.ResponseWriter, r *http.Request) {
		results := []ToDo{
			ToDo{"abc123", "first todo", "this is the first item", 1591559402263, 1591559402263, Open},
			ToDo{"def456", "second todo", "this is the second item", 1591559412263, 1591559412263, Open},
			ToDo{"ghi789", "third todo", "this is the third item", 1591559422263, 1591559422263, Open},
		}
		encoder := json.NewEncoder(w)
		if err := encoder.Encode(results); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/in_progress", func(w http.ResponseWriter, r *http.Request) {
		results := []ToDo{
			ToDo{"abc123", "first todo", "this is the first item", 1591559402263, 1591559402263, InProgress},
		}
		encoder := json.NewEncoder(w)
		if err := encoder.Encode(results); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/closed", func(w http.ResponseWriter, r *http.Request) {
		results := []ToDo{
			ToDo{"abc123", "first todo", "this is the first item", 1591559402263, 1591559402263, Closed},
		}
		encoder := json.NewEncoder(w)
		if err := encoder.Encode(results); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println(http.ListenAndServe(":8080", nil))
}
