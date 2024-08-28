package main

import (
	"fmt"
	"time"
)

type Order struct {
	ID       string
	Date     string
	Products []string
	Price    int
}

type Builder interface {
	Create()
	AddProduct(string)
	GetOrder() Order
}

type BuilderOrder struct {
	order Order
}

func (b *BuilderOrder) Create() {
	b.order.ID = "UUID-123123"
	b.order.Date = time.Now().Format(time.RFC3339)
}

func (b *BuilderOrder) AddProduct(product string) {
	b.order.Products = append(b.order.Products, product)
	b.order.Price += 1
}

func (b *BuilderOrder) GetOrder() Order {
	return b.order
}

func main() {
	builder := &BuilderOrder{}
	builder.Create()
	builder.AddProduct("Product 1")
	builder.AddProduct("Product 2")
	order := builder.GetOrder()

	fmt.Printf("Order: %+v \n", order)
}
