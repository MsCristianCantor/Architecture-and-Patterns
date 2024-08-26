package main

import (
	"fmt"
)

// Printer: Interfaz enfocada en la impresión
type Printer interface {
	Print()
}

// Scanner: Interfaz enfocada en el escaneo
type Scanner interface {
	Scan()
}

// MultiFunctionPrinter: Implementa ambas interfaces
type MultiFunctionPrinter struct{}

func (m MultiFunctionPrinter) Print() {
	fmt.Println("Printing document...")
}

func (m MultiFunctionPrinter) Scan() {
	fmt.Println("Scanning document...")
}

func main() {
	var printer Printer = MultiFunctionPrinter{}
	printer.Print()

	var scanner Scanner = MultiFunctionPrinter{}
	scanner.Scan()
}
