package router

import (
	"fmt"
	"gooop/event"
)

type GenericEventRoute struct {
	event event.Event
}

func (er *GenericEventRoute) Route() {
	fmt.Printf("Routed Generic Event: %s [%T]\n", er.event.GetName(), er.event)
}
