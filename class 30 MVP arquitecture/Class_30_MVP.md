# Class 30: Model-View-Presenter (MVP)

## Objectives
- Understand the **MVP architectural pattern** and how it differs from MVC.  
- Learn the responsibilities of **Model, View, and Presenter**.  
- Explore the benefits of MVP such as **better testability** and **loose coupling**.  
- Implement MVP in Go with a working example.  

---

## What is MVP?
The **Model-View-Presenter (MVP)** is a design pattern derived from MVC, often used in GUI and layered applications.  

- **Model**: Manages data, business rules, and logic.  
- **View**: Displays information and captures user actions. It is **passive**, meaning it doesn’t contain logic.  
- **Presenter**: Handles all logic, retrieves and updates the model, and tells the view what to display.  

💡 The key difference from MVC is that in MVP the **View never directly communicates with the Model** — everything passes through the Presenter.  

---

## UML Diagram

```
+-----------+        +-------------+        +---------+
|   Model   | <----> |  Presenter  | <----> |  View   |
+-----------+        +-------------+        +---------+
```

---

## Conceptual Example
Imagine a login screen:  

- **Model**: Contains the user credentials and validation logic.  
- **View**: Shows the login form and error messages.  
- **Presenter**: Receives user input from the view, asks the model to validate, and updates the view with the result.  

---

## Go Example

```go
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
```

---

## Practical Example
Build a **Task Manager** app:  
- **Model**: Task struct with fields `ID`, `Title`, and `Done`.  
- **View**: Console or Web UI showing the task list.  
- **Presenter**: Handles creating, completing, and listing tasks.  

---

## Additional Resources
- [Martin Fowler on GUI Architectures](https://martinfowler.com/eaaDev/uiArchs.html)  
- [MVP vs MVC vs MVVM](https://www.geeksforgeeks.org/mvc-vs-mvp-vs-mvvm-design-pattern/)  
- Book: *Patterns of Enterprise Application Architecture* by Martin Fowler.  

---

## Optional Challenge
- Extend the login system to support multiple users.  
- Add persistence (store users in a slice or file).  
- Write unit tests for the Presenter.  

---

## Next Steps
In the next class, we will explore **MVVM (Model-View-ViewModel)** and see how it differs from MVC and MVP.  
