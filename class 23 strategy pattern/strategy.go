package main

import "fmt"

// Strategy Interface
type PaymentStrategy interface {
	Pay(amount float64)
}

// Concrete Strategies
type CreditCardPayment struct {
	name   string
	number string
	cvv    string
}

func (c *CreditCardPayment) Pay(amount float64) {
	fmt.Printf("Pagando %.2f usando tarjeta de crédito de %s\n", amount, c.name)
}

type PayPalPayment struct {
	email string
}

func (p *PayPalPayment) Pay(amount float64) {
	fmt.Printf("Pagando %.2f usando PayPal con el correo %s\n", amount, p.email)
}

// Context
type PaymentContext struct {
	strategy PaymentStrategy
}

func (p *PaymentContext) SetStrategy(strategy PaymentStrategy) {
	p.strategy = strategy
}

func (p *PaymentContext) Pay(amount float64) {
	if p.strategy == nil {
		fmt.Println("No se ha configurado un método de pago.")
		return
	}
	p.strategy.Pay(amount)
}

// Main
func main() {
	context := &PaymentContext{}

	// Usando tarjeta de crédito
	creditCard := &CreditCardPayment{name: "Carlos", number: "1234-5678-9876-5432", cvv: "123"}
	context.SetStrategy(creditCard)
	context.Pay(100.0)

	// Usando PayPal
	paypal := &PayPalPayment{email: "carlos@example.com"}
	context.SetStrategy(paypal)
	context.Pay(250.0)
}
