package event

type EventInterface interface {
	GetName() string
	GetType() string
}

type Event struct {
	name string
}

func NewEvent(name string) *Event {
	return &Event{name: name}
}

func (e *Event) GetName() string {
	return e.name
}

func (e *Event) GetType() string {
	return "generic"
}
