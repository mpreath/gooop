package router

import (
	"gooop/event"
	"log"
)

type GenericEventRoute struct {
	event event.Event
}

func (er *GenericEventRoute) Route() {
	log.Printf("Routed Generic Event: %s [%T]\n", er.event.GetName(), er.event)
}
