package main

import "fmt"

type Component interface {
	Operation() string
	Add(Component)
	Remove(Component)
}

type File struct {
	name string
}

func (l *File) Operation() string {
	return l.name
}

func (l *File) Add(c Component) {
	// No hace nada ya que es un file (hoja)
}

func (l *File) Remove(c Component) {
	// No hace nada ya que es un file (hoja)
}

type Folder struct {
	name     string
	children []Component
}

func (c *Folder) Operation() string {
	result := fmt.Sprintf("%s [\n", c.name)
	for _, child := range c.children {
		result += fmt.Sprintf("	%s \n", child.Operation())
	}
	result += "]"
	return result
}

func (c *Folder) Add(component Component) {
	c.children = append(c.children, component)
}

func (c *Folder) Remove(component Component) {
	// Implementar lógica para eliminar el componente
}

func main() {
	file1 := &File{name: "File 1"}
	file2 := &File{name: "File 2"}
	file3 := &File{name: "File 3"}
	folder1 := &Folder{name: "Folder 1"}
	folder2 := &Folder{name: "Folder 2"}
	folder1.Add(file1)
	folder1.Add(file2)
	folder2.Add(file3)
	folder2.Add(folder1)

	fmt.Println(folder2.Operation())
}
