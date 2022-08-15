package router

import (
	"fmt"
	"gooop/event"
)

type DomainEventRoute struct {
	event event.Event
}

func (er *DomainEventRoute) Route() {
	fmt.Printf("Routed Domain Event: %s [%T]\n", er.event.GetName(), er.event)
}
