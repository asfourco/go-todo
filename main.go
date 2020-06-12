package main

import (
	"encoding/json"
	"fmt"
	"go-todo/models"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testTags := []models.Tag{models.Tag{ID: 1, Name: "foo", CreatedAt: "2020-06-12T14:05:26Z"}, models.Tag{ID: 2, Name: "bar", CreatedAt: "2020-06-12T14:05:26Z"}}
		results := []models.ToDo{
			models.ToDo{ID: 123, Title: "first todo", Description: "this is the first item", CreatedAt: "2020-06-12T14:05:26Z", UpdatedAt: "2020-06-12T14:05:26Z", Status: models.Open, Tags: []models.Tag{}},
			models.ToDo{ID: 456, Title: "second todo", Description: "this is the second item", CreatedAt: "2020-06-12T14:05:26Z", UpdatedAt: "2020-06-12T14:05:26Z", Status: models.Closed, Tags: testTags},
			models.ToDo{ID: 789, Title: "third todo", Description: "this is the third item", CreatedAt: "2020-06-12T14:05:26Z", UpdatedAt: "2020-06-12T14:05:26Z", Status: models.InProgress, Tags: []models.Tag{}},
		}
		encoder := json.NewEncoder(w)
		if err := encoder.Encode(results); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/open", func(w http.ResponseWriter, r *http.Request) {
		results := []models.ToDo{
			models.ToDo{ID: 123, Title: "first todo", Description: "this is the first item", CreatedAt: "2020-06-12T14:05:26Z", UpdatedAt: "2020-06-12T14:05:26Z", Status: models.Open, Tags: []models.Tag{}},
		}
		encoder := json.NewEncoder(w)
		if err := encoder.Encode(results); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/in_progress", func(w http.ResponseWriter, r *http.Request) {
		results := []models.ToDo{
			models.ToDo{ID: 123, Title: "first todo", Description: "this is the first item", CreatedAt: "2020-06-12T14:05:26Z", UpdatedAt: "2020-06-12T14:05:26Z", Status: models.InProgress, Tags: []models.Tag{models.Tag{ID: 1, Name: "foo", CreatedAt: "2020-06-12T14:05:26Z"}, models.Tag{ID: 2, Name: "bar", CreatedAt: "2020-06-12T14:05:26Z"}}},
		}
		encoder := json.NewEncoder(w)
		if err := encoder.Encode(results); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/closed", func(w http.ResponseWriter, r *http.Request) {
		results := []models.ToDo{
			models.ToDo{ID: 123, Title: "first todo", Description: "this is the first item", CreatedAt: "2020-06-12T14:05:26Z", UpdatedAt: "2020-06-12T14:05:26Z", Status: models.Closed, Tags: []models.Tag{models.Tag{ID: 1, Name: "foo", CreatedAt: "2020-06-12T14:05:26Z"}, models.Tag{ID: 2, Name: "bar", CreatedAt: "2020-06-12T14:05:26Z"}}},
		}
		encoder := json.NewEncoder(w)
		if err := encoder.Encode(results); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println(http.ListenAndServe(":8080", nil))
}
