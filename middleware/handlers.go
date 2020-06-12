package middleware

/*
Middleware responsible for CRUD operations and handling traffic between the API endpoints and the database.

Based on tutorial from codesource.io - https://codesource.io/build-a-crud-application-in-golang-with-postgresql/
*/

import (
	"encoding/json" // package to encode and decode the json into struct and vice versa
	"fmt"
	"strconv"

	"log"
	"net/http" // used to access the request and response object of the api

	// used to read the environment variable
	// package used to covert string into int type

	"github.com/gorilla/mux" // used to get the params from the route

	_ "github.com/mattn/go-sqlite3" // sqlite3 driver

	"go-todo/models" // models package where ToDo schema is defined
)

// response format
type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

// CreateTodo create a todo entry
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	// set the header to content type x-www-form-urlencoded
	// Allow all origin to handle cors issue
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// create an empty todo
	var todo models.ToDo

	// decode the json request to todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v\n", err)
	}

	newTodo, err := insertTodo(todo)
	if err != nil {
		log.Fatalf("Error in inserting todo %v\n", err)
	}

	err = json.NewEncoder(w).Encode(newTodo)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// GetTodo get a todo
func GetTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// get the todo id from the request params, key is "id"
	params := mux.Vars(r)
	// the id type from string to int
	id, err := strconv.Atoi(params["id"])
	checkErr(err)
	todo, err := getTodo(int64(id))
	// call the getUser function with user id to retrieve a single user
	if err != nil {
		log.Fatalf("Unable to get todo. %v", err)
	}

	// send the response
	err = json.NewEncoder(w).Encode(todo)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// GetAllTodos get all todos
func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// get all the users in the db
	todos, err := getAllTodos()

	if err != nil {
		log.Fatalf("Unable to get all todo. %v", err)
	}

	// send all the users as response
	err = json.NewEncoder(w).Encode(todos)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// UpdateTodo update a todo
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// get the userid from the request params, key is "id"
	params := mux.Vars(r)
	// the id type from string to int
	id, err := strconv.Atoi(params["id"])
	checkErr(err)
	var todo models.ToDo

	// decode the json request to user
	err = json.NewDecoder(r.Body).Decode(&todo)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	newTodo, err := updateTodo(int64(id), todo)
	checkErr(err)

	// send the response
	err = json.NewEncoder(w).Encode(newTodo)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// DeleteTodo delete a todo
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// get the userid from the request params, key is "id"
	params := mux.Vars(r)
	// the id type from string to int
	id, err := strconv.Atoi(params["id"])
	checkErr(err)

	deletedRows, err := deleteTodo(int64(id))
	checkErr(err)

	// format the message string
	msg := fmt.Sprintf("Todo deleted successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// AddTag will add a tag
func AddTag(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// create an empty tag
	var tag models.Tag

	// decode the json request to tag
	err := json.NewDecoder(r.Body).Decode(&tag)
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	newEntry, err := insertTag(tag)
	if err != nil {
		log.Fatalf("Error in inserting new tag. %v\n", err)
	}

	err = json.NewEncoder(w).Encode(newEntry)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// DeleteTag will delete a tag
func DeleteTag(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// get the userid from the request params, key is "id"
	params := mux.Vars(r)
	// the id type from string to int
	id, err := strconv.Atoi(params["id"])
	checkErr(err)

	deletedRows, err := deleteTag(int64(id))
	checkErr(err)
	// format the message string
	msg := fmt.Sprintf("Tag deleted successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// GetTag will get a tag
func GetTag(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// get the todo id from the request params, key is "id"
	params := mux.Vars(r)
	// the id type from string to int
	id, err := strconv.Atoi(params["id"])
	checkErr(err)

	tag, err := getTag(int64(id))
	if err != nil {
		log.Fatalf("Error in retrieving tag. %v\n", err)
	}

	err = json.NewEncoder(w).Encode(tag)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// GetAllTags list all tags
func GetAllTags(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// get all the users in the db
	tags, err := getAllTags()

	if err != nil {
		log.Fatalf("Unable to get all tags. %v", err)
	}

	// send all the users as response
	err = json.NewEncoder(w).Encode(tags)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// AssociateTag will associate a tag with a todo
func AssociateTag(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	// get the todo id from the request params, key is "id"
	params := mux.Vars(r)
	// the id type from string to int
	tagID, err := strconv.Atoi(params["tagID"])
	checkErr(err)
	todoID, err := strconv.Atoi(params["todoID"])
	checkErr(err)

	associationID, err := associateTag(int64(tagID), int64(todoID))
	checkErr(err)
	msg := fmt.Sprintf("Tag associated successfully. todos_tags ID: %v", associationID)

	// format the reponse message
	res := response{
		ID:      int64(associationID),
		Message: msg,
	}

	// send the response
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
