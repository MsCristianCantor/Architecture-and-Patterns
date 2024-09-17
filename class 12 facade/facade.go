package main

import "fmt"

// Subsistemas
type SeatBooking struct{}

func (s *SeatBooking) CheckAvailability() bool {
	fmt.Println("Verificando disponibilidad de asientos...")
	return true
}

type PaymentProcessor struct{}

func (p *PaymentProcessor) ProcessPayment() {
	fmt.Println("Procesando el pago...")
}

type TicketIssuer struct{}

func (t *TicketIssuer) SendTicket() {
	fmt.Println("Enviando ticket al cliente...")
}

// Facade
type CinemaFacade struct {
	seatBooking      *SeatBooking
	paymentProcessor *PaymentProcessor
	ticketIssuer     *TicketIssuer
}

func NewCinemaFacade() *CinemaFacade {
	return &CinemaFacade{
		seatBooking:      &SeatBooking{},
		paymentProcessor: &PaymentProcessor{},
		ticketIssuer:     &TicketIssuer{},
	}
}

func (c *CinemaFacade) BookTicket() {
	if c.seatBooking.CheckAvailability() {
		c.paymentProcessor.ProcessPayment()
		c.ticketIssuer.SendTicket()
	}
}

func main() {
	facade := NewCinemaFacade()
	facade.BookTicket()
}
