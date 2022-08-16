package router

import (
	"log"
)

type DomainEventRoute struct {
	EventRoute
}

func (er *DomainEventRoute) Route() {
	log.Printf("Routed Domain Event: %s [%T]\n", er.event.GetName(), er.event)
}
