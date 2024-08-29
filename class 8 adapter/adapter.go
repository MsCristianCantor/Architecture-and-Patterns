package main

import "fmt"

// TargetUserSummary interfaz que el cliente espera
type TargetUserSummary interface {
	Sumary() string
}

// User estructura con la data original y tomando el papel de un tercero al cual no se le pude acceder a la logica
type User struct {
	Name string
	Age  int
}

func (a *User) GetName() string {
	return a.Name
}

func (a *User) GetAge() int {
	return a.Age
}

// Adapter es la estructura que adapta la interfaz de User a TargetUserSummary
type Adapter struct {
	User *User
}

func (a *Adapter) Request() string {
	return fmt.Sprintf(
		"User Sumary, name: %s, age: %d \n",
		a.User.Name,
		a.User.Age,
	)
}

func main() {
	user := &User{
		Name: "Pepito",
		Age:  30,
	}
	adapter := &Adapter{User: user}

	fmt.Println("Adapter result: ", adapter.Request())
}
