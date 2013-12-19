package task

import (
	"encoding/json"
	"fmt"
	"github.com/gocraft/web"
	"strconv"
	"strings"
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
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Content-Type", "application/json")
	users, _ := getUsers()
	b, _ := json.Marshal(users)
	fmt.Fprint(rw, string(b))

}

func (c *Context) UsersCreate(rw web.ResponseWriter, req *web.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Content-Type", "application/json")
	req.ParseForm()
	name := string(req.PostFormValue("name"))
	email := string(req.PostFormValue("email"))
	if name != "" && email != "" {
		user := User{1, name, email}
		createUser(user)
		fmt.Fprint(rw, ResponseJson{"success": true})
	} else {
		fmt.Fprint(rw, ResponseJson{"success": false})
	}

}

func (c *Context) UsersDelete(rw web.ResponseWriter, req *web.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "DELETE,POST")
	rw.Header().Set("Content-Type", "application/json")
	userId, _ := strconv.Atoi(req.PathParams["id"])
	if deleteUser(userId) {
		fmt.Fprint(rw, ResponseJson{"success": true})
		return
	}
	fmt.Fprint(rw, ResponseJson{"success": false})
}

func (c *Context) UsersUpdate(rw web.ResponseWriter, req *web.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "PUT,POST")
	rw.Header().Set("Content-Type", "application/json")
	userId, _ := strconv.Atoi(req.PathParams["id"])
	email := string(req.PostFormValue("email"))
	name := string(req.PostFormValue("name"))
	user := User{userId, name, email}
	if updateUser(user) {
		fmt.Fprint(rw, ResponseJson{"success": true})
		return
	}
	fmt.Fprint(rw, ResponseJson{"success": false})

}

func (c *Context) UsersAction(rw web.ResponseWriter, req *web.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Content-Type", "application/json")
	req.ParseForm()
	action := string(req.PostFormValue("action"))
	userId, _ := strconv.Atoi(req.PathParams["id"])
	switch strings.ToLower(action) {
	case "delete":
		if deleteUser(userId) {
			fmt.Fprint(rw, ResponseJson{"success": true})
			return
		}

	case "update":
		email := string(req.PostFormValue("email"))
		name := string(req.PostFormValue("name"))
		user := User{userId, name, email}
		if updateUser(user) {
			fmt.Fprint(rw, ResponseJson{"success": true})
			return
		}

	default:

	}
	fmt.Fprint(rw, ResponseJson{"success": false})

}

func (c *Context) CreateDB() {
	createTable()
}
