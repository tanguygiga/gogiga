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

func (impl *TodoTxtImpl) getFilePath() string {
	u, err := user.Current()
	if err != nil {
		fmt.Print(err)
	}
	return u.HomeDir + "/todo-test.txt"
}

// GetAll return a list of Todo
func (impl TodoTxtImpl) GetAll() (list []m.Todo, err error) {
	f, err := os.Open(impl.getFilePath())
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
	f, err := os.Open(impl.getFilePath())
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
	f, err := os.Open(impl.getFilePath())
	if err != nil {
		return err
	}
	s := bufio.NewScanner(f)
	i := 0
	var list []string
	for s.Scan() {
		i++
		if i == id {
			continue
		}
		list = append(list, s.Text()+"\n")
	}
	if err = s.Err(); err != nil {
		return err
	}
	err = impl.writeSortedLines(impl.getFilePath(), list)
	if err != nil {
		return err
	}
	return nil
}

func (impl *TodoTxtImpl) writeSortedLines(file string, lines []string) (err error) {
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
