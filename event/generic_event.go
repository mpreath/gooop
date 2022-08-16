package event

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
