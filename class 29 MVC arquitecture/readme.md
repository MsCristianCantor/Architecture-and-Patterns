# Class 29 – Model-View-Controller (MVC)

## 🎯 Objectives
- Understand the **MVC architectural pattern** and its importance in software design.  
- Learn how MVC separates **concerns** into three main components: Model, View, and Controller.  
- Explore how MVC is applied in real-world applications, especially in web development.  
- Implement a **simple MVC structure in Go**.  

---

## 📖 What is MVC?
**Model-View-Controller (MVC)** is a software architectural pattern that separates an application into three interconnected components:  

1. **Model** → Represents the data and business logic.  
2. **View** → Handles the user interface and presentation.  
3. **Controller** → Acts as an intermediary between Model and View, processing input and coordinating responses.  

👉 MVC promotes **separation of concerns**, making applications more maintainable, scalable, and testable.  

---

## 🗂 UML Diagram

+-------------+ +---------------+ +-------------+
| Model | <----> | Controller | <----> | View |
| (Data/Logic)| | (Input Logic) | | (UI Layer) |
+-------------+ +---------------+ +-------------+

- The **Controller** updates the **Model** based on user input.  
- The **Model** notifies the **View** of changes.  
- The **View** displays data from the **Model**.  

---

## 💡 Conceptual Example
Imagine a **Library System**:  
- **Model** → Holds book data (title, author, availability).  
- **View** → Displays a list of available books.  
- **Controller** → Handles user actions like “borrow book” or “return book.”  

This separation ensures that if the **UI changes** (e.g., web → mobile), the **Model and Controller** remain the same.  

---

## 🖥 Golang Example (Simple MVC)

```go
package main

import "fmt"

// Model
type Book struct {
    Title  string
    Author string
}

// View
type BookView struct{}

func (bv *BookView) ShowBookDetails(book Book) {
    fmt.Printf("Book: %s by %s\n", book.Title, book.Author)
}

// Controller
type BookController struct {
    model Book
    view  BookView
}

func (c *BookController) SetBookTitle(title string) {
    c.model.Title = title
}

func (c *BookController) SetBookAuthor(author string) {
    c.model.Author = author
}

func (c *BookController) UpdateView() {
    c.view.ShowBookDetails(c.model)
}

// Main
func main() {
    model := Book{Title: "The Go Programming Language", Author: "Alan Donovan"}
    view := BookView{}
    controller := BookController{model: model, view: view}

    controller.UpdateView()

    // Update data via controller
    controller.SetBookTitle("Design Patterns in Go")
    controller.SetBookAuthor("John Doe")
    controller.UpdateView()
}

```
---

## 🏋️ Practical Example

**Task:**

Build a **Student Management System** using MVC:

* **Model** → Store student information (name, ID, grade).

* **View** → Display student details.

* **Controller** → Allow updating and retrieving student data.

👉 Bonus: Add multiple students and show how the View handles lists.

---

## 📚 Additional Resources

* [Wikipedia - MVC](https://es.wikipedia.org/wiki/Modelo%E2%80%93vista%E2%80%93controlador)
* [Go Patterns – Architectural](https://github.com/tmrts/go-patterns)
* “Head First Design Patterns” – Architectural patterns section.

---

## 🎯 Optional Challenge

* Extend the **Student Management System** into a **mini CLI app** where the user can input commands (add student, show student, update grade).

* Add persistence by saving and loading data to a file.

---

## 🔜 What’s Next?

Next, we will explore **Model-View-Presenter (MVP)** and **Model-View-ViewModel (MVVM)** as alternative architectural patterns, and compare them with MVC.