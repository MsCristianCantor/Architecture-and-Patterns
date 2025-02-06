package main

import "fmt"

// Mediator interface
type ChatMediator interface {
	SendMessage(msg string, user User)
	AddUser(user User)
}

// Concrete Mediator
type ChatRoom struct {
	users []User
}

func (c *ChatRoom) SendMessage(msg string, user User) {
	for _, u := range c.users {
		if u != user {
			u.ReceiveMessage(msg)
		}
	}
}

func (c *ChatRoom) AddUser(user User) {
	c.users = append(c.users, user)
}

// Colleague interface
type User interface {
	SendMessage(msg string)
	ReceiveMessage(msg string)
}

// Concrete Colleague
type ChatUser struct {
	name     string
	mediator ChatMediator
}

func (u *ChatUser) SendMessage(msg string) {
	fmt.Printf("%s envía: %s\n", u.name, msg)
	u.mediator.SendMessage(msg, u)
}

func (u *ChatUser) ReceiveMessage(msg string) {
	fmt.Printf("%s recibe: %s\n", u.name, msg)
}

func main() {
	chatRoom := &ChatRoom{}

	user1 := &ChatUser{name: "Juan", mediator: chatRoom}
	user2 := &ChatUser{name: "Maria", mediator: chatRoom}
	user3 := &ChatUser{name: "Carlos", mediator: chatRoom}

	chatRoom.AddUser(user1)
	chatRoom.AddUser(user2)
	chatRoom.AddUser(user3)

	user1.SendMessage("Hola a todos!")
}
