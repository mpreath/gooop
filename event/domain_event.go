package event

type DomainEvent struct {
	Event
}

func NewDomainEvent(name string) *DomainEvent {
	return &DomainEvent{Event: Event{name: name}}
}

func (e *DomainEvent) GetType() string {
	return "domain"
}
