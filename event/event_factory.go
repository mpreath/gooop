package event

func GenerateEvent(eventName string, eventType string) Event {
	if eventType == "domain" {
		return &DomainEvent{name: eventName}
	} else {
		return &GenericEvent{name: eventName}
	}
}
