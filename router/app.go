package router

import (
	"gogiga/handler"
	"gogiga/stringutil"
	"net/http"
)

// App register all the Handler of the application and their routes
type App struct {
	TodoHandler *handler.TodoHandler
}

func (h *App) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = stringutil.ShiftPath(req.URL.Path)
	switch head {
	case "todos":
		h.TodoHandler.GetAll(res, req)
	case "todo":
		h.TodoHandler.ServeHTTP(res, req)
	default:
		http.Error(res, "Not Found", http.StatusNotFound)
	}
}
