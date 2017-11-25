package dao

import (
	"bufio"
	"fmt"
	m "gogiga/model"
	"os"
	"os/user"
)

// TodoTxtImpl text implementation of Todo
type TodoTxtImpl struct {
}

func (impl *TodoTxtImpl) getFile() (*os.File, error) {
	u, err := user.Current()
	if err != nil {
		fmt.Print(err)
	}
	f, err := os.Open(u.HomeDir + "/todo.txt")
	return f, err
}

// GetAll return a list of Todo
func (impl TodoTxtImpl) GetAll() (list []m.Todo, err error) {
	f, err := impl.getFile()
	if err != nil {
		return nil, err
	}
	s := bufio.NewScanner(f)
	i := 0
	for s.Scan() {
		var t m.Todo
		i++
		t.ID = i
		t.Task = s.Text()
		list = append(list, t)
	}
	if err = s.Err(); err != nil {
		return nil, err
	}

	return list, nil
}

// Get return a Todo
func (impl TodoTxtImpl) Get(id int) (t m.Todo, err error) {
	f, err := impl.getFile()
	if err != nil {
		return t, err
	}
	s := bufio.NewScanner(f)
	i := 0
	for s.Scan() {
		i++
		if i == id {
			t.ID = i
			t.Task = s.Text()
			break
		}
	}
	if err = s.Err(); err != nil {
		return t, err
	}
	return t, nil
}

// Create a Todo
func (impl TodoTxtImpl) Create(t *m.Todo) error {
	fmt.Println("Create(t *m.Todo) : Not yet implemented !")
	return nil
}

// Update a Todo
func (impl TodoTxtImpl) Update(t *m.Todo) error {
	fmt.Println("Update(t *m.Todo) : Not yet implemented !")
	return nil
}

// Delete a Todo
func (impl TodoTxtImpl) Delete(id int) error {
	fmt.Println("Delete(id int) : Not yet implemented !")
	return nil
}
