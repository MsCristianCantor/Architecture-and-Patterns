package main

import "fmt"

// Handler define la interfaz para los manejadores en la cadena
type Handler interface {
	SetNext(handler Handler)
	HandleRequest(request string)
}

// BaseHandler proporciona una implementación básica de SetNext
type BaseHandler struct {
	next Handler
}

func (h *BaseHandler) SetNext(handler Handler) {
	h.next = handler
}

func (h *BaseHandler) HandleNext(request string) {
	if h.next != nil {
		h.next.HandleRequest(request)
	}
}

// LoggerHandler es un handler que solo registra la solicitud
type LoggerHandler struct {
	BaseHandler
}

func (h *LoggerHandler) HandleRequest(request string) {
	fmt.Println("Logger: Registrando la solicitud:", request)
	h.HandleNext(request)
}

// AuthHandler verifica si la solicitud está autenticada
type AuthHandler struct {
	BaseHandler
	isAuthenticated bool
}

func (h *AuthHandler) HandleRequest(request string) {
	if h.isAuthenticated {
		fmt.Println("AuthHandler: Solicitud autenticada.")
		h.HandleNext(request)
	} else {
		fmt.Println("AuthHandler: Solicitud no autenticada. Deteniendo la cadena.")
	}
}

// PermissionHandler verifica si la solicitud tiene los permisos adecuados
type PermissionHandler struct {
	BaseHandler
	hasPermission bool
}

func (h *PermissionHandler) HandleRequest(request string) {
	if h.hasPermission {
		fmt.Println("PermissionHandler: Permisos validados.")
		h.HandleNext(request)
	} else {
		fmt.Println("PermissionHandler: Sin permisos suficientes. Deteniendo la cadena.")
	}
}

func main() {
	// Crear handlers
	logger := &LoggerHandler{}
	auth := &AuthHandler{isAuthenticated: true}
	permission := &PermissionHandler{hasPermission: true}

	// Configurar la cadena
	logger.SetNext(auth)
	auth.SetNext(permission)

	// Enviar una solicitud a través de la cadena
	fmt.Println("Primera solicitud:")
	logger.HandleRequest("Operacion 1")

	// Cambiar permisos para probar un caso fallido
	fmt.Println("\nSegunda solicitud (sin permisos):")
	permission.hasPermission = false
	logger.HandleRequest("Operacion 2")
}
