package router

import (
	"gooop/event"
	"log"
	"sync"
)

type EventRouter struct {
	wg sync.WaitGroup
}

func NewEventRouter() EventRouter {
	return EventRouter{}
}

func (er *EventRouter) RouteEvent(e event.EventInterface) {
	log.Printf("Routing Event: %s [%T]\n", e.GetName(), e)

	r := er.getRoute(e)
	er.wg.Add(1)
	go func() {
		defer er.wg.Done()
		r.Route()
	}()
}

func (er *EventRouter) getRoute(e event.EventInterface) Route {

	// Get route based on event type
	switch v := e.(type) {
	case *event.DomainEvent:
		return &DomainEventRoute{event: v}
	default:
		return &GenericEventRoute{event: v}
	}
}

func (er *EventRouter) Wait() {
	er.wg.Wait()
}
