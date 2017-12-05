package handler

import (
	"encoding/json"
	"fmt"
	"gogiga/dao"
	"gogiga/model"
	"gogiga/stringutil"
	"log"
	"net/http"
	"strconv"
)

// TodoHandler handle requests for Todo
type TodoHandler struct {
	dao dao.TodoDao
}

// NewTodoHandler creating handler with given implementation
func NewTodoHandler() *TodoHandler {
	h := &TodoHandler{}
	h.dao = dao.TodoDaoFactory("txt")
	return h
}

// GetAll get all the todo
func (h *TodoHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	listTodo, err := h.dao.GetAll()
	if err != nil {
		log.Fatal(err)
		return
	}
	json, _ := json.Marshal(listTodo)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", json)
}

func (h *TodoHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = stringutil.ShiftPath(req.URL.Path)
	id, err := strconv.Atoi(head)
	if err != nil {
		http.Error(res, fmt.Sprintf("Invalid todo id %q", head), http.StatusBadRequest)
		return
	}
	switch req.Method {
	case "GET":
		h.get(res, req, id)
	case "POST":
		h.create(res, req)
	case "PUT":
		h.update(res, req, id)
	case "DELETE":
		h.delete(res, req, id)
	default:
		http.Error(res, req.Method+" is not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *TodoHandler) get(w http.ResponseWriter, r *http.Request, id int) {
	fmt.Printf("GET todo ID is >>>>> %d\n", id)
	todo, err := h.dao.Get(id)
	json, _ := json.Marshal(todo)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", json)
}

func (h *TodoHandler) create(w http.ResponseWriter, r *http.Request) {
	t := model.Todo{}
	json.NewDecoder(r.Body).Decode(&t)
	err := h.dao.Create(&t)
	if err != nil {
		log.Fatal(err)
	}
	json, _ := json.Marshal(t)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", json)
}

func (h *TodoHandler) update(w http.ResponseWriter, r *http.Request, id int) {
	t := model.Todo{}
	json.NewDecoder(r.Body).Decode(&t)
	err := h.dao.Update(&t)
	if err != nil {
		log.Fatal(err)
	}
	json, _ := json.Marshal(t)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", json)
}

func (h *TodoHandler) delete(w http.ResponseWriter, r *http.Request, id int) {
	err := h.dao.Delete(id)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s %d", "Removed User", id)
}
