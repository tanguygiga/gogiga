package dao

import (
	"bufio"
	"fmt"
	m "gogiga/model"
	"os"
	"os/user"
	"sort"
)

// TodoTxtImpl text implementation of Todo
type TodoTxtImpl struct {
}

func getFilePath() string {
	u, err := user.Current()
	if err != nil {
		fmt.Print(err)
	}
	return u.HomeDir + "/todo-test.txt"
}

// GetAll return a list of Todo
func (impl TodoTxtImpl) GetAll() (listTodo []m.Todo, err error) {
	f, err := os.Open(getFilePath())
	if err != nil {
		return nil, err
	}
	s := bufio.NewScanner(f)
	for i := 1; s.Scan(); i++ {
		var t m.Todo
		t.ID = i
		t.Task = s.Text()
		listTodo = append(listTodo, t)
	}
	if err = s.Err(); err != nil {
		return nil, err
	}

	return listTodo, nil
}

// Get return a Todo
func (impl TodoTxtImpl) Get(id int) (t m.Todo, err error) {
	f, err := os.Open(getFilePath())
	if err != nil {
		return t, err
	}
	s := bufio.NewScanner(f)
	for i := 1; s.Scan(); i++ {
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
	f, err := os.Open(getFilePath())
	if err != nil {
		return err
	}
	s := bufio.NewScanner(f)
	var list []string
	for s.Scan() {
		list = append(list, s.Text()+"\n")
	}
	list = append(list, t.Task+"\n")
	if err = s.Err(); err != nil {
		return err
	}
	err = writeSortedLines(getFilePath(), list)
	if err != nil {
		return err
	}
	return nil
}

// Update a Todo
func (impl TodoTxtImpl) Update(t *m.Todo) error {
	f, err := os.Open(getFilePath())
	if err != nil {
		return err
	}
	s := bufio.NewScanner(f)
	var list []string
	for i := 1; s.Scan(); i++ {
		task := s.Text()
		if i == t.ID {
			task = t.Task
		}
		list = append(list, task+"\n")
	}
	if err = s.Err(); err != nil {
		return err
	}
	err = writeSortedLines(getFilePath(), list)
	if err != nil {
		return err
	}
	return nil
}

// Delete a Todo
func (impl TodoTxtImpl) Delete(id int) error {
	f, err := os.Open(getFilePath())
	if err != nil {
		return err
	}
	s := bufio.NewScanner(f)
	var list []string
	for i := 1; s.Scan(); i++ {
		if i == id {
			continue
		}
		list = append(list, s.Text()+"\n")
	}
	if err = s.Err(); err != nil {
		return err
	}
	err = writeSortedLines(getFilePath(), list)
	if err != nil {
		return err
	}
	return nil
}

func writeSortedLines(file string, lines []string) error {
	sort.Strings(lines)
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	defer w.Flush()
	for _, line := range lines {
		if "\n" == line {
			continue
		}
		_, err := w.WriteString(line)
		if err != nil {
			return err
		}
	}
	return nil
}
