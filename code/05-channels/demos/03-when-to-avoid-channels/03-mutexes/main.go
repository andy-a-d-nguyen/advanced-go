package main

import "sync"

func main() {

}

type User struct {
	Username string
}

var users = make([]User, 0)
var m = sync.Mutex{}

func getAll() []User {
	m.Lock()
	defer m.Unlock()
	return users
}

func add(username string) User {
	m.Lock()
	defer m.Unlock()

	u := User{username}
	users = append(users, u)
	return u
}
