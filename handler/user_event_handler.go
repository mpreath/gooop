package handler

import "log"

type UserEventHandler struct {
	event Event
}

func (h *UserEventHandler) Handle() {
	log.Printf("handling user event: %T", h.event)
}
