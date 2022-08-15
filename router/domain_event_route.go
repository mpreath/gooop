package router

import (
	"gooop/event"
	"log"
)

type DomainEventRoute struct {
	event event.Event
}

func (er *DomainEventRoute) Route() {
	log.Printf("Routed Domain Event: %s [%T]\n", er.event.GetName(), er.event)
}
