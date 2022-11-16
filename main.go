package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.Use(loggingMiddleware)

	r.HandleFunc("/todos", listTodosHandler).Methods(http.MethodGet)
	r.HandleFunc("/todos/{id:[0-9]+}", getTodoHandler).Methods(http.MethodGet)
	r.HandleFunc("/todos", createTodoHandler).Methods(http.MethodPost)
	r.HandleFunc("/todos/{id:[0-9]+}", updateTodoHandler).Methods(http.MethodPatch)
	// r.HandleFunc("/todos", deleteTodoHandler).Methods(http.MethodDelete)
	// r.HandleFunc("/todos", deleteAllTodosHandler).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":8080", r))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
