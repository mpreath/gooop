package handler

import "log"

type PropertyEventHandler struct {
	event Event
}

func (h *PropertyEventHandler) Handle() {
	log.Printf("handling property event: %T", h.event)
}
