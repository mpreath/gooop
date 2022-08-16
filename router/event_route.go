package router

import (
	"gooop/event"
	"log"
)

type EventRouteInterface interface {
	Route()
}

type EventRoute struct {
	event event.EventInterface
}

func (er *EventRoute) Route() {
	log.Printf("Routed Generic Event: %s [%T]\n", er.event.GetName(), er.event)
}
