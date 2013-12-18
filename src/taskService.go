package task

import (
	"fmt"
	"github.com/vmihailenco/pg"
)

type User struct {
	Id    int
	Name  string
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

func connectDB() *pg.DB {
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "python",
		Database: "myapp",
	})
	return db
}

func createTable() {
	db := connectDB()
	defer db.Close()
	res, err := db.Exec("CREATE TABLE users(id serial PRIMARY KEY , name varchar(500), email varchar(500))")
	fmt.Println(res, err)
}

func deleteUser(id int) bool {
	db := connectDB()
	defer db.Close()
	_, err := db.Exec("DELETE FROM users WHERE id = ?", id)
	if err == nil {
		return true
	}
	return false
}

func createUser(user User) {
	db := connectDB()
	defer db.Close()
	res, err := db.Exec("INSERT INTO users(name,email) VALUES(?, ?)", user.Name, user.Email)
	fmt.Println(res.Affected(), err)
}

func updateUser(user User) bool {
	db := connectDB()
	defer db.Close()
	_, err := db.Exec("UPDATE users SET name=?, email = ? WHERE id=?", user.Name, user.Email, user.Id)
	if err == nil {
		return true
	}
	return false

}

func getUsers() ([]*User, error) {
	db := connectDB()
	defer db.Close()
	users := &Users{}
	_, err := db.Query(users, "SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	return users.Values, nil
}

func viewUsers(users []*User) {
	for _, element := range users {
		fmt.Println(element.Name)
		fmt.Println(element.Email)
	}
}
