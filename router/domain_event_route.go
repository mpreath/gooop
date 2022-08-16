package router

import (
	"gooop/event"
	"log"
)

type DomainEventRoute struct {
	event event.EventInterface
}

func (er *DomainEventRoute) Route() {
	log.Printf("Routed Domain Event: %s [%T]\n", er.event.GetName(), er.event)
}
