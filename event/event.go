package event

type EventInterface interface {
	GetName() string
	GetType() string
	SetHandled(bool)
	GetHandled() bool
}

type Event struct {
	name    string
	handled bool
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

func (e *Event) SetHandled(v bool) {
	e.handled = v
}

func (e *Event) GetHandled() bool {
	return e.handled
}
