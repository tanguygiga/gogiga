package controller

import (
	"fmt"
	"gogiga/dao"
	m "gogiga/model"
)

// TodoController controller
type TodoController struct {
	dao dao.TodoDao
}

// NewTodoController creating controller with given implementation
func NewTodoController(impl string) *TodoController {
	tc := &TodoController{}
	tc.dao = dao.TodoDaoFactory(impl)
	return tc
}

// GetAll print all todo
func (tc TodoController) GetAll() {
	listTodo, err := tc.dao.GetAll()
	if err != nil {
		return
	}
	fmt.Println("--- todo.txt ---")
	for i := range listTodo {
		fmt.Println(listTodo[i].ID, "\t", listTodo[i].Task)
	}
}

// Get print a Todo
func (tc TodoController) Get(id int) {
	todo, err := tc.dao.Get(id)
	if err != nil {
		return
	}
	fmt.Printf("--- Ligne %d ---\n", id)
	fmt.Println(todo.ID, "\t", todo.Task)
}

// Set Create or Update a Todo and print the list
func (tc TodoController) Set(id int, task string) {
	todo, err := tc.dao.Get(id)
	if err != nil {
		return
	}
	if todo.ID == 0 {
		// Create
		t := m.Todo{}
		t.ID, t.Task = id, task
		if err = tc.dao.Create(&t); err != nil {
			return
		}
		fmt.Printf("--- Création d'une ligne : %s\n", task)
	} else {
		// Update
		t := m.Todo{}
		t.ID, t.Task = id, task
		if err = tc.dao.Update(&t); err != nil {
			return
		}
		fmt.Printf("--- Mise à jour de la ligne %d ---\n", id)
		fmt.Println(id, "\t"+task)
	}
	tc.GetAll()
}

// Delete a Todo and print the remaining list
func (tc TodoController) Delete(id int) {
	err := tc.dao.Delete(id)
	if err != nil {
		return
	}
	fmt.Printf("--- Suppression de la ligne %d ---\n", id)
	tc.GetAll()
}
