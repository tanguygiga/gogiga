package model

// Todo ID: int, Task: string
type Todo struct {
	ID   int
	Task string
}

// NewTodo constructor
func NewTodo(id int, task string) *Todo {
	return &Todo{ID: id, Task: task}
}
