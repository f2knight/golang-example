package task

import (
	"github.com/gocraft/web"
    "fmt"
    "encoding/json"
)

type Context struct {
	test string
}

func (c *Context) Root(rw web.ResponseWriter, req *web.Request) {
    fmt.Fprint(rw, "Welcome")
}

func (c *Context) UsersList(rw web.ResponseWriter, req *web.Request) {
    users,_ := getUsers()
    for _,element := range users {
    	b, _ := json.Marshal(element)
    	fmt.Fprint(rw, string(b))
    }    
}

func (c *Context) UsersCreate(rw web.ResponseWriter, req *web.Request) {
	req.ParseForm()
	name := string(req.PostFormValue("name"))
	email := string(req.PostFormValue("email"))
    user := User{1,name,email}    
    createUser(user) 
}

