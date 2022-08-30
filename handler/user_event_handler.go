package handler

import (
	"fmt"
	"log"
)

var validActions = []string{"updated", "created", "deleted"}

type UserEventHandler struct {
	event Event
}

func (h *UserEventHandler) Handle() error {
	log.Printf("handling user event: %T", h.event)

	if !validateAction(h.event.GetAttributes()["action"]) {
		return fmt.Errorf("invalid action in event")
	}

	log.Println("user event handled")
	return nil
}

func validateAction(action string) bool {
	validEvent := false
	for _, validAction := range validActions {
		if action == validAction {
			validEvent = true
			break
		}
	}
	return validEvent
}
