package middleware

import (
	"database/sql"
	"fmt"
	"go-todo/models"
	"log"
)

func initialiseToDo(db *sql.DB) error {
	// create todos table
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS todos (id INTEGER PRIMARY KEY, title TEXT, description TEXT, createdAt TIMESTAMP default (strftime('%s', 'now')), updatedAt TIMESTAMP DEFAULT (strftime('%s', 'now')), status INTEGER)")
	checkErr(err)

	_, err = statement.Exec()
	return err
}

func initialiseTag(db *sql.DB) error {
	// create tags table
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS tags (id INTEGER PRIMARY KEY, name STRING, createdAt TIMESTAMP default (strftime('%s', 'now')))")
	checkErr(err)

	_, err = statement.Exec()
	return err
}

func initialiseTodoTag(db *sql.DB) error {
	// create todos table
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS todos_tags (id INTEGER PRIMARY KEY, todo_id INTEGER, tag_id INTEGER)")
	checkErr(err)

	_, err = statement.Exec()
	return err
}

func checkOrCreateTables(db *sql.DB) error {
	if err := initialiseToDo(db); err != nil {
		return err
	}

	if err := initialiseTag(db); err != nil {
		return err
	}

	if err := initialiseTodoTag(db); err != nil {
		return err
	}

	return nil
}

// create connection with sqlite db
func createConnection() *sql.DB {
	// Open the connection
	db, err := sql.Open("sqlite3", "../db/todo.db")
	checkErr(err)

	// check the connection
	err = db.Ping()
	checkErr(err)

	err = checkOrCreateTables(db)
	checkErr(err)
	// return the connection
	return db
}

//------------------------- handler functions ----------------

func insertTodo(todo models.ToDo) (models.ToDo, error) {
	db := createConnection()
	defer db.Close()

	statement, err := db.Prepare("INSERT INTO todos (title, description, status) VALUES (?, ?, 0)")
	checkErr(err)

	response, err := statement.Exec(todo.Title, todo.Description)
	checkErr(err)

	id, err := response.LastInsertId()
	if err != nil {
		log.Fatalf("Unable to execute the query. %v\n", err)
	}

	fmt.Printf("Inserted a single record %v\n", id)

	newEntry, err := getTodo(id)
	checkErr(err)
	// return the inserted id
	return newEntry, err
}

func getTodo(id int64) (models.ToDo, error) {
	db := createConnection()
	defer db.Close()

	var todo models.ToDo
	row := db.QueryRow("SELECT * FROM todos WHERE id=?", id)

	err := row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt, &todo.Status)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return todo, nil
	case nil:
		tags, err := getTagsOfTodo(id)
		if err != nil {
			fmt.Printf("No tags for todo id: %v\n", id)
		}
		todo.Tags = tags

		return todo, err
	default:
		log.Fatalf("getTodo: Unable to scan the row. %v\n", err)
		return todo, err
	}

}

func getAllTodos() ([]models.ToDo, error) {
	db := createConnection()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM todos")
	checkErr(err)

	var todos []models.ToDo
	var todo models.ToDo
	for rows.Next() {
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt, &todo.Status); err != nil {
			log.Fatalf("getAllTodos: Unable to scan the row. %v\n", err)
		}
		todos = append(todos, todo)
	}
	return todos, err
}

func updateTodo(id int64, todo models.ToDo) (models.ToDo, error) {
	db := createConnection()
	defer db.Close()

	fmt.Printf("received %v\n", todo)
	statement, err := db.Prepare("UPDATE todos SET title=?, description=?, status=?, updatedAt=strftime('%s', 'now') WHERE id=?")
	checkErr(err)

	response, err := statement.Exec(todo.Title, todo.Description, todo.Status, id)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v\n", err)
	}

	rowsAffected, err := response.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v\n", err)
	}
	fmt.Printf("Total rows/record affected %v\n", rowsAffected)

	updatedEntry, err := getTodo(id)
	checkErr(err)

	return updatedEntry, err
}

func deleteTodo(id int64) (int64, error) {
	db := createConnection()
	defer db.Close()

	statement, err := db.Prepare("DELETE FROM todos WHERE id=?")
	checkErr(err)

	response, err := statement.Exec(id)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v\n", err)
	}

	rowsAffected, err := response.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v\n", err)
	}

	fmt.Printf("Total rows/record affected %v\n", rowsAffected)

	return rowsAffected, err
}

func getTagsOfTodo(id int64) ([]models.Tag, error) {
	db := createConnection()
	defer db.Close()

	statement, err := db.Prepare("SELECT t.id, t.name, t.createdAt FROM todos_tags jt JOIN tags t on t.id = jt.tag_id WHERE jt.todo_id=?")
	checkErr(err)

	rows, err := statement.Query(id)
	checkErr(err)
	var tags []models.Tag
	var tag models.Tag
	for rows.Next() {
		if err := rows.Scan(&tag.ID, &tag.Name, &tag.CreatedAt); err != nil {
			log.Fatalf("getTagsOfTodo: Unable to scan the row. %v\n", err)
		}
		tags = append(tags, tag)
	}
	return tags, err
}

func insertTag(tag models.Tag) (models.Tag, error) {
	db := createConnection()
	defer db.Close()

	statement, err := db.Prepare("INSERT INTO tags (name) VALUES (?)")
	checkErr(err)

	response, err := statement.Exec(tag.Name)
	checkErr(err)

	id, err := response.LastInsertId()
	if err != nil {
		log.Fatalf("Unable to execute the query. %v\n", err)
	}

	fmt.Printf("Inserted a single record %v\n", id)
	newEntry, err := getTag(id)
	checkErr(err)
	// return the inserted id
	return newEntry, err
}

func getTag(id int64) (models.Tag, error) {
	db := createConnection()
	defer db.Close()

	row := db.QueryRow("SELECT * FROM tags WHERE id=?", id)

	var tag models.Tag
	err := row.Scan(&tag.ID, &tag.Name, &tag.CreatedAt)
	if err != nil {
		fmt.Printf("getTag: Unable to Scan row. %v\n", err)
	}
	return tag, err
}

func deleteTag(id int64) (int64, error) {
	db := createConnection()
	defer db.Close()

	statement, err := db.Prepare("DELETE FROM tags WHERE id=?")
	checkErr(err)

	response, err := statement.Exec(id)
	if err != nil {
		log.Fatalf("deleteTag: Unable to execute the query. %v\n", err)
	}

	rowsAffected, err := response.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v\n", err)
	}

	fmt.Printf("Total rows/record affected %v\n", rowsAffected)

	return rowsAffected, err
}

func associateTag(tagID int64, todoID int64) (int64, error) {
	db := createConnection()
	defer db.Close()

	statement, err := db.Prepare("INSERT INTO todos_tags (tag_id, todo_id) VALUES (?, ?)")
	checkErr(err)

	response, err := statement.Exec(tagID, todoID)
	checkErr(err)

	id, err := response.LastInsertId()
	if err != nil {
		log.Fatalf("Unable to associate tag to todo. %v\n", err)
	}

	// return the inserted id
	return id, err
}

func getAllTags() ([]models.Tag, error) {
	db := createConnection()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM tags")
	checkErr(err)

	var tags []models.Tag
	var tag models.Tag
	for rows.Next() {
		if err := rows.Scan(&tag.ID, &tag.Name, &tag.CreatedAt); err != nil {
			log.Fatalf("getAllTags: Unable to scan the row. %v\n", err)
		}
		tags = append(tags, tag)
	}
	return tags, err
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
