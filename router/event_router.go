package router

import (
	"fmt"
	"gooop/event"
)

func RouteEvent(e event.Event) {
	fmt.Printf("Routing Event: %s [%T]\n", e.GetName(), e)

	GetRoute(e).Route()
}

func GetRoute(e event.Event) Route {

	// Get route based on event type
	switch v := e.(type) {
	case *event.DomainEvent:
		return &DomainEventRoute{event: v}
	default:
		return &GenericEventRoute{event: v}
	}
}
