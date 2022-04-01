package main

import "fmt"

func main() {

}

type User struct {
	Username string
}

func getAll() []User {
	msg := getAllMsg{resultCh: make(chan []User)}
	getAllCh <- msg
	<-msg.resultCh
}

func add(username string) User {
	msg := addMsg{
		username: username,
		resultCh: make(chan User),
		errorCh:  make(chan error),
	}
	addCh <- msg
	select {
	case u := <-msg.resultCh:
		fmt.Println("User created")
	case err := <-msg.errorCh:
		fmt.Println(err)
	}
}

func managerUsers() {
	users := []User{}
	for {
		select {
		case msg := <-getAllCh:
			msg.resultCh <- users
		case msg := <-addCh:
			u := User{msg.username}
			users = append(users, u)
			msg.resultCh <- u
		}
	}
}

type getAllMsg struct {
	resultCh chan []User
}
type addMsg struct {
	username string
	resultCh chan User
	errorCh  chan error
}

var (
	getAllCh = make(chan getAllMsg)
	addCh    = make(chan addMsg)
)
