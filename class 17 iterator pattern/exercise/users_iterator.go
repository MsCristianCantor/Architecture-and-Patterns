package main

import "fmt"

// Iterator define los métodos de iteración
type Iterator interface {
	HasNext() bool
	Next() interface{}
}

// Collection define la interfaz para obtener un iterador
type Collection interface {
	CreateIterator() Iterator
}

type UsersCollection struct {
	users []User
}

// User define la estructura de un usuario
type User struct {
	name string
	age  int
}

func (c *UsersCollection) CreateIterator() Iterator {
	return &ConcreteIterator{
		collection: c,
		index:      0,
	}
}

type ConcreteIterator struct {
	collection *UsersCollection
	index      int
}

func (i *ConcreteIterator) HasNext() bool {
	return i.index < len(i.collection.users)
}

func (i *ConcreteIterator) Next() interface{} {
	if i.HasNext() {
		item := i.collection.users[i.index]
		i.index++
		return item
	}
	return nil
}

func main() {
	collection := &UsersCollection{
		users: []User{
			{"John", 25},
			{"Jane", 30},
			{"Doe", 35},
		},
	}
	iterator := collection.CreateIterator()

	for iterator.HasNext() {
		fmt.Println(iterator.Next())
	}
}
