package task

import (
	"github.com/gocraft/web"
    "fmt"
)

type Users struct {
    Values []*User
}

type Context struct {
    HelloCount int
}

func (c *Context) Root(rw web.ResponseWriter, req *web.Request) {
    fmt.Fprint(rw, "Welcome")
}

