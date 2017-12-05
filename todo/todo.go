package main

import (
	"gogiga/handler"
	"gogiga/router"
	"net/http"
)

func main() {
	app := &router.App{
		TodoHandler: handler.NewTodoHandler(),
	}

	http.ListenAndServe(":8080", app)
}
