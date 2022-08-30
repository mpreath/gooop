package handler

import "log"

type PropertyEventHandler struct {
	event Event
}

func (h *PropertyEventHandler) Handle() error {
	log.Printf("handling property event: %T", h.event)

	log.Printf("property event handled")
	return nil
}
