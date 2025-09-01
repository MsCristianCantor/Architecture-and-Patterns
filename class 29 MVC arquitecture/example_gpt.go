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
