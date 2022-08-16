package event

func GenerateEvent(eventName string, eventType string) EventInterface {
	if eventType == "domain" {
		return &DomainEvent{Event: Event{name: eventName}}
	} else {
		return &GenericEvent{Event: Event{name: eventName}}
	}
}
