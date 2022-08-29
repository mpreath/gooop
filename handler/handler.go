package handler

import "log"

type Event interface {
	GetAttributes() map[string]string
	GetName() string
}

type BaseEvent struct {
	Name       string
	Attributes map[string]string
}

type UserEvent struct {
	BaseEvent
}

func (e *BaseEvent) GetName() string {
	return e.Name
}

func (e *BaseEvent) GetAttributes() map[string]string {
	return e.Attributes
}

type EventHandler interface {
	Handle()
}

type UserEventHandler struct {
	event Event
}

func (h *UserEventHandler) Handle() {
	log.Printf("handling user event: %T", h.event)
}

func HandleEvent(event Event) {
	switch event.(type) {
	case *UserEvent:
		handler := UserEventHandler{event: event}
		handler.Handle()
	default:
		log.Printf("no handler for: %T", event)
	}
}
