package handler

import (
	"gooop/event"
	"log"
)

type LogHandler struct {
	BaseHandler
}

func (h *LogHandler) Execute(e event.EventInterface) {
	log.Printf("Logged Event: %s [%T]\n", e.GetName(), e)
	if h.next != nil {
		h.next.Execute(e)
	}
}
