package event

import "fmt"

func New(eventName string, eventType string) (EventInterface, error) {
	switch eventType {
	case "domain":
		return NewDomainEvent(eventName), nil
	case "event":
		return NewEvent(eventName), nil
	default:
		return nil, fmt.Errorf("unknown type \"%s\" for event \"%s\"", eventType, eventName)
	}
}
