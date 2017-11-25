package main

import (
	"gogiga/controller"
	"net/http"
)

func main() {

	tc := controller.NewTodoController("txt")
	tc.GetAll()
	tc.Get(5)
	//tc.Delete(5)

	http.ListenAndServe(":8088", handler)
}
