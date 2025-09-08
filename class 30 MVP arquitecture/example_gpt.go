package main

import "fmt"

// Model
type User struct {
	Username string
	Password string
}

func (u *User) Validate() bool {
	return u.Username == "admin" && u.Password == "1234"
}

// View
type LoginView interface {
	ShowSuccess()
	ShowError()
}

// Presenter
type LoginPresenter struct {
	view LoginView
}

func (p *LoginPresenter) Login(u *User) {
	if u.Validate() {
		p.view.ShowSuccess()
	} else {
		p.view.ShowError()
	}
}

// Concrete View
type ConsoleLoginView struct{}

func (c *ConsoleLoginView) ShowSuccess() {
	fmt.Println("✅ Login successful!")
}

func (c *ConsoleLoginView) ShowError() {
	fmt.Println("❌ Login failed. Try again.")
}

func main() {
	view := &ConsoleLoginView{}
	presenter := &LoginPresenter{view}

	user := &User{Username: "admin", Password: "1234"}
	presenter.Login(user)

	user2 := &User{Username: "john", Password: "wrong"}
	presenter.Login(user2)
}
