package main

import (
	"fmt"
)

// Database: Interface abstracta para bases de datos
type Database interface {
	Save(data string)
}

// MySQLDatabase: Implementa la interfaz Database para MySQL
type MySQLDatabase struct{}

func (db MySQLDatabase) Save(data string) {
	fmt.Printf("Saving '%s' to MySQL database\n", data)
}

// UserService: Depende de la abstracción Database, no de la implementación concreta
type UserService struct {
	db Database
}

func (u *UserService) SaveUser(data string) {
	u.db.Save(data)
}

func main() {
	mysqlDB := MySQLDatabase{}
	userService := UserService{db: mysqlDB}
	userService.SaveUser("John Doe")
}
