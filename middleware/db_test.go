package middleware

import (
	"go-todo/models"
	"testing"
)

func TestConnection(t *testing.T) {
	db := createConnection()
	defer db.Close()

	if err := db.Ping(); err != nil {
		t.Error("Cannot connect to database", err)
	}
}

func TestDBInsertTodo(t *testing.T) {
	var todo models.ToDo
	todo.Title = "testing insert todo"
	todo.Description = "Some description"
	todo.Status = models.Open

	newEntry, err := insertTodo(todo)
	if err != nil {
		t.Error("Error in adding new Todo", err)
	}
	if newEntry.IsEmpty() {
		t.Error("Did not create todo")
	}

	if newEntry.Title != todo.Title {
		t.Error("New entry title does not match original")
	}

	if newEntry.Status != models.Open {
		t.Error("New entry status is not Open")
	}
}

func TestDBGetTodo(t *testing.T) {
	var todo models.ToDo
	todo.Title = "testing get todo"
	todo.Description = "Some description"
	todo.Status = 0

	newEntry, err := insertTodo(todo)
	if err != nil {
		t.Error("Error in adding new Todo", err)
	}

	retrievedTodo, err := getTodo(int64(newEntry.ID))

	if err != nil {
		t.Error("Error in retrieving Todo", err)
	}

	if retrievedTodo.ID != newEntry.ID {
		t.Error("Retrieved Todo ID does not match request")
	}

	if retrievedTodo.Title != todo.Title {
		t.Error("Retrieved Todo title does not match original")
	}

	if retrievedTodo.Description != todo.Description {
		t.Error("Retrieved Todo description does not match original")
	}

	if retrievedTodo.Status != models.Open {
		t.Error("Retrieved Todo status is not 'Open' ")
	}
}

func TestDBGetAllTodos(t *testing.T) {
	todos, err := getAllTodos()
	if err != nil {
		t.Error("Error fetching todos", err)
	}
	if len(todos) < 1 {
		t.Error("No Todo rows returned")
	}
}

func TestDBUpdateTodo(t *testing.T) {
	todos, err := getAllTodos()
	if err != nil {
		t.Error("Error fetching todos", err)
	}
	todo := todos[0]

	todo.Status = models.InProgress

	updatedEntry, err := updateTodo(todo.ID, todo)
	if err != nil {
		t.Error("Error updating todo", err)
	}

	if updatedEntry.Status != models.InProgress {
		t.Error("Todo was not updated", todo, updatedEntry)
	}
}

func TestDBDeleteTodo(t *testing.T) {
	todos, err := getAllTodos()
	if err != nil {
		t.Error("Error fetching todos", err)
	}
	todo := todos[0]

	_, err = deleteTodo(todo.ID)
	if err != nil {
		t.Error("Todo was not deleted", err)
	}

	response, _ := getTodo(todo.ID)

	if !response.IsEmpty() {
		t.Error("Was able to fetch the original todo", response, todo)
	}
}

func TestDBinsertTag(t *testing.T) {
	var tag models.Tag
	tag.Name = "testing tag"
	newEntry, err := insertTag(tag)
	if err != nil {
		t.Error("Error inserting tag", err)
	}
	if newEntry.IsEmpty() {
		t.Error("Returned Tag is empty")
	}
}

func TestDBDeleteTag(t *testing.T) {
	tags, _ := getAllTags()
	tag := tags[0]
	if _, err := deleteTag(tag.ID); err != nil {
		t.Error("Error deleting tag", err)
	}
}

func TestDBGetAllTags(t *testing.T) {
	tags, err := getAllTags()
	if err != nil {
		t.Error("Error in fetching all tags", err)
	}
	if len(tags) < 1 {
		t.Error("Expected to have at least one tag record")
	}
}

func TestDBAssociateTagWithTodo(t *testing.T) {
	todos, err := getAllTodos()
	if err != nil {
		t.Error("Error in fetching all todos", err)
	}

	todo := todos[0]

	var newTag models.Tag
	newTag.Name = "association"
	response, err := insertTag(newTag)
	if err != nil {
		t.Error("Error creating new Tag", err)
	}

	if _, err = associateTag(response.ID, todo.ID); err != nil {
		t.Error("Error associating tag", err)
	}

	updatedTodo, err := getTodo(todo.ID)
	if err != nil {
		t.Error("Error retrieving updated todo", err)
	}

	if updatedTodo.ID != todo.ID {
		t.Error("Retrieved Todo does not match original")
	}

	if len(updatedTodo.Tags) < 1 {
		t.Error("Retrieved Todo has no tags associated", updatedTodo.Tags)
	}

	if updatedTodo.Tags[0].Name != newTag.Name {
		t.Error("Retrieved Todo tag does not match entry")
	}

}
