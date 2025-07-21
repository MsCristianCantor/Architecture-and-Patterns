package main

import "fmt"

// Element interface
type Element interface {
	Accept(Visitor)
}

// Visitor interface
type Visitor interface {
	VisitBook(*Book)
	VisitMagazine(*Magazine)
}

// Concrete Elements
type Book struct {
	Title string
}

func (b *Book) Accept(v Visitor) {
	v.VisitBook(b)
}

type Magazine struct {
	Title string
}

func (m *Magazine) Accept(v Visitor) {
	v.VisitMagazine(m)
}

// Concrete Visitor
type PrintVisitor struct{}

func (v *PrintVisitor) VisitBook(b *Book) {
	fmt.Println("Libro:", b.Title)
}

func (v *PrintVisitor) VisitMagazine(m *Magazine) {
	fmt.Println("Revista:", m.Title)
}

func main() {
	elements := []Element{
		&Book{Title: "El principito"},
		&Magazine{Title: "National Geographic"},
	}

	visitor := &PrintVisitor{}
	for _, el := range elements {
		el.Accept(visitor)
	}
}
