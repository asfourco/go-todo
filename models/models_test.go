package models

import (
	"testing"
)

func TestStatus(t *testing.T) {
	var s ToDoStatus

	s = Open
	if s.String() != "Open" {
		t.Error("Todo status is not 'Open'")
	}
	s = InProgress
	if s.String() != "In Progress" {
		t.Error("Todo status is not 'In Progress'")
	}

	s = Closed
	if s.String() != "Closed" {
		t.Error("Todo status is not 'Closed'")
	}

	s = 10

	if s.String() != "Unknown" {
		t.Error("Todo status is not 'Unkown'")
	}
}

func TestTodoStruct(t *testing.T) {
	todo := ToDo{12, "test", "test description", "2020-06-12T14:05:26Z", "2020-06-12T14:05:26Z", 0, []Tag{}}

	if todo.IsEmpty() {
		t.Error("New todo is empty")
	}

	if todo.Status != Open {
		t.Error("New todo status is not Open")
	}
}

func TestChangeTodoStatus(t *testing.T) {
	todo := ToDo{12, "test", "test description", "2020-06-12T14:05:26Z", "2020-06-12T14:05:26Z", 0, []Tag{}}

	todo.Status = InProgress
	if todo.Status != InProgress {
		t.Error("Todo status not changed to In Progress")
	}

	todo.Status = Closed
	if todo.Status != Closed {
		t.Error("Todo status not changed to Closed")
	}
}

func TestTagStruct(t *testing.T) {
	tag := Tag{1, "tag", "2020-06-12T14:05:26Z"}

	if tag.IsEmpty() {
		t.Error("New tag is empty")
	}

	if tag.Name != "tag" {
		t.Error("Tag name does not match expected")
	}
}

func TestAssociateTagToTodo(t *testing.T) {
	tags := []Tag{
		Tag{1, "tagOne", "2020-06-12T14:05:26Z"},
		Tag{2, "tagTwo", "2020-06-12T14:05:26Z"},
	}

	todo := ToDo{12, "test", "test description", "2020-06-12T14:05:26Z", "2020-06-12T14:05:26Z", 0, tags}

	if len(todo.Tags) != 2 {
		t.Error("Todo tags is not of expected length", todo)
	}

	if todo.Tags[0].Name != "tagOne" {
		t.Error("First tag does not match expcted tag", todo)
	}
}
