package main

import (
	"errors"
	"fmt"
	"regexp"
)

// Handler define la interfaz para los manejadores en la cadena
type Handler interface {
	SetNext(handler Handler)
	HandleRequest(request string) error
}

// BaseHandler proporciona una implementación básica de SetNext
type BaseHandler struct {
	next Handler
}

func (h *BaseHandler) SetNext(handler Handler) {
	h.next = handler
}

func (h *BaseHandler) HandleNext(request string) error {
	if h.next != nil {
		return h.next.HandleRequest(request)
	}
	return nil
}

type EmptyValidationHandler struct {
	BaseHandler
}

func (h *EmptyValidationHandler) HandleRequest(request string) error {
	if len(request) > 0 {
		fmt.Println("la cadena no esta vacia")
		return h.HandleNext(request)
	}

	return errors.New("la cadena esta vacia")
}

type MiniumLenValidationHandler struct {
	BaseHandler
}

func (h *MiniumLenValidationHandler) HandleRequest(request string) error {
	if len(request) > 5 {
		fmt.Println("la cadena es lo soficientemente larga")
		return h.HandleNext(request)
	}

	return errors.New("la cadena no es lo soficientemente larga")
}

type CharectersValidationHandler struct {
	BaseHandler
}

func (h *CharectersValidationHandler) HandleRequest(request string) error {
	regex := regexp.MustCompile(`[^\w\s]`)

	if !regex.MatchString(request) {
		fmt.Println("la cadena no contiene caracteres especiales")
		return h.HandleNext(request)
	}

	return errors.New("la cadena contiene caracteres especiales")
}

func main() {
	emptyValidation := &EmptyValidationHandler{}
	miniumLenValidation := &MiniumLenValidationHandler{}
	charectersValidation := &CharectersValidationHandler{}

	emptyValidation.SetNext(miniumLenValidation)
	miniumLenValidation.SetNext(charectersValidation)

	fmt.Println("Primera solicitud:")
	err := emptyValidation.HandleRequest("test 1")
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("-----------")
	fmt.Println("Segunda solicitud vacia:")
	err = emptyValidation.HandleRequest("")
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("-----------")
	fmt.Println("Tercera solicitud cadena corta:")
	err = emptyValidation.HandleRequest("tes")
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("-----------")
	fmt.Println("Cuarta solicitud cadena con caracteres especiales:")
	err = emptyValidation.HandleRequest("test 4 _?:!@#$")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
