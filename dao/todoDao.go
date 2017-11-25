package dao

import m "gogiga/model"

// TodoDao Data Access Object for Todo
type TodoDao interface {
	GetAll() ([]m.Todo, error)
	Get(id int) (m.Todo, error)
	Create(t *m.Todo) error
	Update(t *m.Todo) error
	Delete(id int) error
}
