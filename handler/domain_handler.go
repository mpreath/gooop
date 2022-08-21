package handler

import (
	"gooop/event"
	"log"
)

type DomainHandler struct {
	BaseHandler
}

func NewDomainHandler() *DomainHandler {
	return &DomainHandler{BaseHandler: BaseHandler{next: nil}}
}

func (h *DomainHandler) Execute(e event.EventInterface) {
	if !e.GetHandled() {
		log.Printf("Handled Domain Event: %s [%T]\n", e.GetName(), e)
		e.SetHandled(true)
	} else {
		log.Printf("Event Already Was Handled: %s [%T]\n", e.GetName(), e)
	}
	if h.next != nil {
		h.next.Execute(e)
	}
}
