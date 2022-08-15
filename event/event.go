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

type GenericEvent struct {
	name string
}

func NewGenericEvent() GenericEvent {
	e := GenericEvent{name: ""}
	return e
}

func (e *GenericEvent) SetName(name string) {
	e.name = name
}

func (e *GenericEvent) GetName() string {
	return e.name
}
