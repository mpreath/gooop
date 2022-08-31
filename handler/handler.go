package handler

import (
	"fmt"
	"log"
)

type EventHandler interface {
	Handle() error
}

func HandleEvent(event Event) error {
	var handler EventHandler

	switch event.(type) {
	case *UserEvent:
		handler = &UserEventHandler{event: event}
	case *PropertyEvent:
		handler = &PropertyEventHandler{event: event}
	default:
		log.Printf("no handler for: %T", event)
		return nil
	}

	err := handler.Handle()
	if err != nil {
		return fmt.Errorf("HandleEvent(): %v", err)
	}
	return nil
}
