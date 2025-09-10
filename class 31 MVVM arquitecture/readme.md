# Class 31: Model-View-ViewModel (MVVM)

## Objectives
- Understand the **MVVM architectural pattern** and its main components.  
- Learn how MVVM separates responsibilities to improve scalability and testability.  
- Explore how MVVM is commonly used in UI-driven applications.  
- Implement a **simple MVVM example in Golang**.  

---

## What is MVVM?
**MVVM (Model-View-ViewModel)** is a software architectural pattern that helps separate the user interface (UI) from business logic.  

It introduces an intermediate layer, the **ViewModel**, which acts as a bridge between the **View** and the **Model**.  

- **Model:** Represents the data and business logic.  
- **View:** The UI, which displays the data to the user.  
- **ViewModel:** Exposes data from the Model in a format that the View can easily use. It also handles commands/events.  

This separation allows:  
- Easier **unit testing** (logic can be tested without the UI).  
- Improved **maintainability**.  
- Reusable **ViewModels** across multiple Views.  

---

## UML Diagram
```
+-----------+        +-------------+        +---------+
|   View    | <----> |  ViewModel  | <----> |  Model  |
+-----------+        +-------------+        +---------+
     |                      |                     |
     | User actions         | Data binding        | Business logic
     v                      v                     v
```

---

## Conceptual Example
Imagine a **Login Screen**:  
- **Model:** Holds user credentials and validates them.  
- **View:** The login form UI.  
- **ViewModel:** Connects the input fields in the UI with the validation logic in the Model.  

The **View** only shows fields and buttons.  
The **ViewModel** processes input (e.g., check if email is valid) and passes data to the Model.  
The **Model** contains the real logic (authentication).  

---

## Golang Example
```go
package main

import "fmt"

// Model
type User struct {
	Email    string
	Password string
}

func (u *User) Authenticate() bool {
	// Simulate authentication
	return u.Email == "admin@example.com" && u.Password == "1234"
}

// ViewModel
type LoginViewModel struct {
	User *User
}

func (vm *LoginViewModel) Login(email, password string) string {
	vm.User.Email = email
	vm.User.Password = password

	if vm.User.Authenticate() {
		return "✅ Login successful"
	}
	return "❌ Login failed"
}

// View
type LoginView struct {
	ViewModel *LoginViewModel
}

func (v *LoginView) DisplayLogin(email, password string) {
	result := v.ViewModel.Login(email, password)
	fmt.Println(result)
}

func main() {
	// Initialize components
	user := &User{}
	viewModel := &LoginViewModel{User: user}
	view := &LoginView{ViewModel: viewModel}

	// Simulate user interaction
	view.DisplayLogin("admin@example.com", "1234")
	view.DisplayLogin("user@example.com", "wrongpass")
}
```

---

## Practical Example
👉 Implement a **To-Do List app** with MVVM:  
- **Model:** Task (ID, title, completed status).  
- **ViewModel:** Exposes methods like `AddTask`, `RemoveTask`, and `ListTasks`.  
- **View:** Displays tasks and interacts with the user through the ViewModel.  

---

## Additional Resources
- [MVVM Explained (Microsoft Docs)](https://learn.microsoft.com/en-us/dotnet/architecture/mvvm/)  
- [MVVM in practice](https://www.geeksforgeeks.org/mvvm-design-pattern/)  
- [MVVM vs MVC vs MVP](https://www.baeldung.com/cs/mvc-vs-mvp-vs-mvvm)  

---

## Optional Challenge
- Extend the **To-Do List app** to allow marking tasks as completed and filtering tasks.  
- Write unit tests for the **ViewModel**, ensuring it works independently from the View.  

---

## Next steps
In the next class, we will move on to **Layered Architecture**, where we will structure applications using multiple layers (presentation, business, and data access).  
