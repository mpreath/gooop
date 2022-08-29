package handler

import "log"

type EventHandler interface {
	Handle()
}

func HandleEvent(event Event) {
	var handler EventHandler

	switch event.(type) {
	case *UserEvent:
		handler = &UserEventHandler{event: event}
	case *PropertyEvent:
		handler = &PropertyEventHandler{event: event}
	default:
		log.Printf("no handler for: %T", event)
		return
	}

	handler.Handle()
}