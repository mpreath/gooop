package event

type EventInterface interface {
	GetName() string
}

type Event struct {
	name string
}

func (e *Event) GetName() string {
	return e.name
}
