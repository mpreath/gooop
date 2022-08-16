package event

func GenerateEvent(eventName string, eventType string) EventInterface {
	if eventType == "domain" {
		return NewDomainEvent(eventName)
	} else {
		return NewEvent(eventName)
	}
}
