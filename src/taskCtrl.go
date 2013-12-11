package task

import (
	"github.com/gocraft/web"
    "fmt"
    "encoding/json"
)

type Context struct {
	test string
}


type Response map[string]interface{}

func (r Response) String() (s string) {
        b, err := json.Marshal(r)
        if err != nil {
                s = ""
                return
        }
        s = string(b)
        return
}

func (c *Context) Root(rw web.ResponseWriter, req *web.Request) {
    fmt.Fprint(rw, "Welcome")
}

func (c *Context) UsersList(rw web.ResponseWriter, req *web.Request) {    
    var result = ""
    rw.Header().Set("Access-Control-Allow-Origin", "*")
    rw.Header().Set("Content-Type", "application/json")
    users,_ := getUsers()   
    for _,element := range users {
    	b, _ := json.Marshal(element)
        result += string(b)
    }    
    fmt.Fprint(rw, Response{"content": result})
}

func (c *Context) UsersCreate(rw web.ResponseWriter, req *web.Request) {
	req.ParseForm()
	name := string(req.PostFormValue("name"))
	email := string(req.PostFormValue("email"))
    user := User{1,name,email}   
    createUser(user) 
    
    rw.Header().Set("Access-Control-Allow-Origin", "*")
    rw.Header().Set("Content-Type", "application/json")
    fmt.Fprint(rw, Response{"success": true})
}

