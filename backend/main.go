package main

import (
	"log"
	"net/http"
	"simple-todo/backend/handlers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	})

	r.HandleFunc("/todos", handlers.GetTodos).Methods("GET")
	r.HandleFunc("/todos", handlers.AddTodo).Methods("POST")
	r.HandleFunc("/todos/{id}", handlers.UpdateTodo).Methods("PUT")
	r.HandleFunc("/todos/{id}", handlers.DeleteTodo).Methods("DELETE")
	r.HandleFunc("/", handlers.Test).Methods("GET")

	log.Println("Starting server on :8081")
	if err := http.ListenAndServe(":8081", c.Handler(r)); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
