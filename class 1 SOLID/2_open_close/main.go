package main

import (
	"fmt"
)

// PaymentProcessor: Interface que define un procesador de pagos
type PaymentProcessor interface {
	ProcessPayment(amount float64)
}

// PayPal: Implementa el procesador de pagos para PayPal
type PayPal struct{}

func (p PayPal) ProcessPayment(amount float64) {
	fmt.Printf("Processing PayPal payment of %f\n", amount)
}

// CreditCard: Implementa el procesador de pagos para tarjeta de crédito
type CreditCard struct{}

func (c CreditCard) ProcessPayment(amount float64) {
	fmt.Printf("Processing Credit Card payment of %f\n", amount)
}

func main() {
	var processor PaymentProcessor = PayPal{}
	processor.ProcessPayment(100.0)

	processor = CreditCard{}
	processor.ProcessPayment(200.0)
}
