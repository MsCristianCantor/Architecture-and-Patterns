package main

import "fmt"

type INotify interface {
	getMessage() string
}

type BaseNotify struct{}

func (d *BaseNotify) getMessage() string {
	return "Esta es la base de la notificacion"
}

type UrgentNotify struct {
	notify INotify
}

func (d *UrgentNotify) getMessage() string {
	return fmt.Sprintf("[Urgente] %s", d.notify.getMessage())
}

type ConfidencialNotify struct {
	notify INotify
}

func (d *ConfidencialNotify) getMessage() string {
	return fmt.Sprintf("[Confidencial] %s", d.notify.getMessage())
}

func main() {
	notify := &BaseNotify{}

	urgentNotify := &UrgentNotify{
		notify: notify,
	}

	confidenciaAndUrgentlNotify := &ConfidencialNotify{
		notify: urgentNotify,
	}

	fmt.Printf("Notificacion: %s\n", confidenciaAndUrgentlNotify.getMessage())
}
