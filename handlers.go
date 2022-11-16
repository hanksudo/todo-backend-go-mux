package main

import (
	"encoding/json"
	"io"
	"net/http"
)

var todoRespository = NewTodoRepository()

func listTodosHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	todos := todoRespository.GetAll()
	json.NewEncoder(w).Encode(todos)
}

func createTodoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	b, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	todo := &Todo{}

	err = json.Unmarshal(b, todo)
	if err != nil {
		panic(err)
	}

	todoRespository.Create(todo)
	todo.SetURL(r)
	json.NewEncoder(w).Encode(todo)
}
