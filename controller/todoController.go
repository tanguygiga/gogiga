package controller

import (
	"fmt"
	"gogiga/dao"
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
	fmt.Println(todo.ID, "\t", todo.Task)
}
