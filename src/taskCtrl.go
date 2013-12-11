package task

import (
	"github.com/gocraft/web"
    "fmt"
    "encoding/json"
    //"strings"
    "strconv"
)

type Context struct {
	test string
}


type ResponseJson map[string]interface{}

func (r ResponseJson) String() (s string) {
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
    var jsonMap = make(ResponseJson)  
    rw.Header().Set("Access-Control-Allow-Origin", "*")
    rw.Header().Set("Content-Type", "application/json")
    users,_ := getUsers()   
    for _,element := range users {
        var jsonMapItem = make(ResponseJson)
        jsonMapItem["id"] = strconv.Itoa(element.Id)
        jsonMapItem["name"] = element.Name
        jsonMapItem["email"] = element.Email
        jsonMap[strconv.Itoa(element.Id)] = jsonMapItem
    }    
    fmt.Fprint(rw, jsonMap)
}

func (c *Context) UsersCreate(rw web.ResponseWriter, req *web.Request) {
	req.ParseForm()
	name := string(req.PostFormValue("name"))
	email := string(req.PostFormValue("email"))
    user := User{1,name,email}   
    createUser(user) 
    
    rw.Header().Set("Access-Control-Allow-Origin", "*")
    rw.Header().Set("Content-Type", "application/json")
    fmt.Fprint(rw, ResponseJson{"success": true})
}

