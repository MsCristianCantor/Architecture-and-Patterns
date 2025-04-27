package main

import "fmt"

// Compression Interface
type CompressionStrategy interface {
	Compress(fileName string)
}

// Concrete Strategies
type ZipCompression struct{}

func (s *ZipCompression) Compress(fileName string) {
	fmt.Printf("Archivo %s comprimido usando ZIP\n", fileName)
}

// Concrete Strategies
type RarCompression struct{}

func (s *RarCompression) Compress(fileName string) {
	fmt.Printf("Archivo %s comprimido usando RAR\n", fileName)
}

// Context
type CompressionContext struct {
	strategy CompressionStrategy
}

func (c *CompressionContext) SetStrategy(strategy CompressionStrategy) {
	c.strategy = strategy
}

func (c *CompressionContext) Compress(fileName string) {
	if c.strategy == nil {
		fmt.Println("No se ha configurado un metodo de compresion.")
		return
	}
	c.strategy.Compress(fileName)
}

// Main
func main() {
	context := &CompressionContext{}
	context.Compress("archivo1.txt")

	// Usando compresion ZIP
	zip := &ZipCompression{}
	context.SetStrategy(zip)
	context.Compress("archivo1.txt")

	// Usando compresion RAR
	rar := &RarCompression{}
	context.SetStrategy(rar)
	context.Compress("archivo2.txt")
}
