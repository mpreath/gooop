package event

type DomainEvent struct {
	name string
}

func NewDomainEvent() DomainEvent {
	e := DomainEvent{name: ""}
	return e
}

func (e *DomainEvent) SetName(name string) {
	e.name = name
}

func (e *DomainEvent) GetName() string {
	return e.name
}
