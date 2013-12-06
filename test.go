package main

import (
    "fmt"
    "github.com/vmihailenco/pg"
)

type User struct {
    Id int
    Name   string
    Email string
}

type Users struct {
    Values []*User
}

func (f *Users) New() interface{} {
    u := &User{}
    f.Values = append(f.Values, u)
    return u
}

func createTable(db *pg.DB) {
    res, err :=db.Exec("CREATE TABLE users(id serial PRIMARY KEY , name varchar(500), email varchar(500))")
    fmt.Println(res.Affected(), err)
}

func createUser(db *pg.DB , user User){
    res, err := db.Exec("INSERT INTO users(name,email) VALUES(?, ?)", user.Name, user.Email)
    fmt.Println(res.Affected(), err)
}

func getUsers(db *pg.DB) ([]*User, error) {
    users := &Users{}
    _, err := db.Query(users,"SELECT * FROM users")
    if err != nil {
        return nil, err
    }
    return users.Values, nil
}

func delUser(db *pg.DB) {
    
}


func createUsers(db *pg.DB) {
    user := User{1,"Na","tinanguyen@gmail.com"}
    createUser(db,user)
}

func viewUsers(users []*User) {
    for _,element := range users {
        fmt.Println(element)
    }
}

func main() {
    db := pg.Connect(&pg.Options{
        User: "postgres",
        Password : "python",
        Database : "myapp",
    })
    defer db.Close()
    createUsers(db)
    //createTable(db);
    //createUser(db)
    users,err := getUsers(db)
    fmt.Println(err)
    viewUsers(users)
}