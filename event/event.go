package event

type EventInterface interface {
	GetName() string
	GetType() string
}

type Event struct {
	name string
}

func (e *Event) GetName() string {
	return e.name
}

func (e *Event) GetType() string {
	return "generic"
}
