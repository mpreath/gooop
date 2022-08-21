package handler

import (
	"fmt"
	"gooop/event"
	"log"
)

type Handler interface {
	Execute(event.EventInterface)
	SetNext(Handler)
}

func New(handler_type string) (Handler, error) {
	switch handler_type {
	case "domain":
		return &DomainHandler{BaseHandler: BaseHandler{next: nil}}, nil
	case "log":
		return &LogHandler{BaseHandler: BaseHandler{next: nil}}, nil
	default:
		return nil, fmt.Errorf("unknown handler type \"%s\"", handler_type)
	}
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
