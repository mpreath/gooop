package handler

import (
	"gooop/event"
	"log"
)

type Handler interface {
	Execute(event.EventInterface)
	SetNext(Handler)
}

type BaseHandler struct {
	next Handler
}

func (h *BaseHandler) Execute(e event.EventInterface) {

	if !e.GetHandled() {
		log.Printf("Execute method not defined.\n")
	} else {
		log.Printf("Event was already handled.\n")
	}
}

func (h *BaseHandler) SetNext(next Handler) {
	h.next = next
}
