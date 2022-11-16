package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var todoRespository = NewTodoRepository()

func listTodosHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	todos := todoRespository.GetAll()
	json.NewEncoder(w).Encode(todos)
}

func getTodoHandler(w http.ResponseWriter, r *http.Request) {
	idVar := mux.Vars(r)["id"]

	id, err := strconv.Atoi(idVar)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	todo := todoRespository.Get(id)
	json.NewEncoder(w).Encode(todo)
}

func createTodoHandler(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	todo := &Todo{}

	err = json.Unmarshal(b, todo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	todoRespository.Create(todo)
	todo.SetURL(r)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

func updateTodoHandler(w http.ResponseWriter, r *http.Request) {
	idVar := mux.Vars(r)["id"]

	id, err := strconv.Atoi(idVar)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	existsTodo := todoRespository.Get(id)
	if existsTodo == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("not found"))
		return
	}

	b, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	updateTodo := &Todo{}
	err = json.Unmarshal(b, updateTodo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	existsTodo.Completed = updateTodo.Completed
	existsTodo.Order = updateTodo.Order
	existsTodo.Title = updateTodo.Title
	todoRespository.Update(existsTodo)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(existsTodo)
}

func deleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	idVar := mux.Vars(r)["id"]

	id, err := strconv.Atoi(idVar)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = todoRespository.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func deleteAllTodosHandler(w http.ResponseWriter, r *http.Request) {
	todoRespository.DeleteAll()
	w.WriteHeader(http.StatusNoContent)
}
