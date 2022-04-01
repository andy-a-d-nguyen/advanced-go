package main

func main() {

}

type User struct {
	Username string
}

var users = make([]User, 0)

func getAll() []User {
	return users
}

func add(username string) User {
	u := User{username}
	users = append(users, u)
	return u
}
