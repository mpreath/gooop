package chain

import (
	"gooop/event"
	"log"
)

type LogHandler struct {
	BaseHandler
}

func (h *LogHandler) Execute(e event.EventInterface) {
	log.Printf("Logged Event: %s [%T]\n", e.GetName(), e)
	h.next.Execute(e)
}
