package main

import (
	"fmt"
)

// Report: Maneja la generación de reportes
type Report struct {
	Title string
	Body  string
}

// SaveToFile: Maneja la lógica de guardar un archivo
func SaveToFile(filename string, content string) {
	fmt.Printf("Saving report to %s\n", filename)
	// Lógica para guardar el archivo
}

func main() {
	report := Report{Title: "Monthly Report", Body: "This is the report body"}
	SaveToFile("report.txt", report.Body)
}
