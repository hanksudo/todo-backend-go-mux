package main

import "fmt"

type TodoRespository interface {
	Create(todo *Todo)
	Get(id int) *Todo
	GetAll() []*Todo
	Update(*Todo) error
	Delete(id int) error
	DeleteAll()
}

type InMemoryTodoRepository struct {
	Todos  []*Todo
	nextID int
}

func NewTodoRepository() TodoRespository {
	return &InMemoryTodoRepository{Todos: []*Todo{}, nextID: 1}
}

func (r *InMemoryTodoRepository) Create(todo *Todo) {
	todo.ID = r.nextID
	r.Todos = append(r.Todos, todo)
	r.nextID++
}

func (r *InMemoryTodoRepository) Get(id int) *Todo {
	for _, t := range r.Todos {
		if t.ID == id {
			return t
		}
	}
	return nil
}

func (r *InMemoryTodoRepository) GetAll() []*Todo {
	return r.Todos
}

func (r *InMemoryTodoRepository) Update(todo *Todo) error {
	for i, t := range r.Todos {
		if t.ID == todo.ID {
			r.Todos[i] = todo
			return nil
		}
	}
	return fmt.Errorf("todo not found")
}

func (r *InMemoryTodoRepository) Delete(id int) error {
	for i, t := range r.Todos {
		if t.ID == id {
			r.Todos = append(r.Todos[:i], r.Todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("todo not found")
}

func (r *InMemoryTodoRepository) DeleteAll() {
	r.Todos = []*Todo{}
}
