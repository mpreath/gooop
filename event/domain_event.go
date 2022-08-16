package event

type DomainEvent struct {
	Event
}

func (e *DomainEvent) GetType() string {
	return "domain"
}
