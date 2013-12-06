package main

import (
    "github.com/gocraft/web"
    "net/http"
    "./src"
)



func main() {
    router := web.New(task.Context{}).                  
        Middleware(web.LoggerMiddleware).           
        Middleware(web.ShowErrorsMiddleware).       
        //Middleware((*task.Context).SetHelloCount).  
        //Get("/users", (*task.Context).UsersList)   
        //Post("/users", (*task.Context).UsersCreate)  
        //Put("/users/:id", (*task.Context).UsersUpdate)
        //Delete("/users/:id", (*task.Context).UsersDelete)
        //Patch("/users/:id", (*task.Context).UsersUpdate)
        Get("/", (*task.Context).Root)               
    http.ListenAndServe("localhost:3001", router)   
}
