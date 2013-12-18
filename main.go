package main

import (
	"./src"
	"github.com/gocraft/web"
	"net/http"
)

func main() {
	router := web.New(task.Context{}).
		Middleware(web.LoggerMiddleware).
		Middleware(web.ShowErrorsMiddleware).
		Get("/users", (*task.Context).UsersList).
		Post("/users", (*task.Context).UsersCreate).
		Post("/users/:id:\\d+", (*task.Context).UsersAction). //for Del and Update use Ajax
		Delete("/users/:id:\\d+", (*task.Context).UsersDelete).
		Put("/users/:id:\\d+", (*task.Context).UsersUpdate).
		Get("/", (*task.Context).Root)

	http.ListenAndServe("localhost:3001", router)
}
