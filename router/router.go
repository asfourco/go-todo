package router

import (
	"go-todo/middleware"

	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	// Todo routes
	router.HandleFunc("/api/todo/{id}", middleware.GetTodo).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/todo", middleware.GetAllTodos).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/todo", middleware.CreateTodo).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/todo/{id}", middleware.UpdateTodo).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/todo/{id}", middleware.DeleteTodo).Methods("DELETE", "OPTIONS")

	// Tag routes
	router.HandleFunc("/api/tag/{id}", middleware.GetTag).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/tag", middleware.GetAllTags).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/tag", middleware.AddTag).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/tag/{id}", middleware.DeleteTag).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/tag/todo/{tagID}/{todoID}", middleware.AssociateTag).Methods("POST", "OPTIONS")

	return router
}
