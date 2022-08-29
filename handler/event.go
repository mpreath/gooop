package handler

type Event interface {
	GetAttributes() map[string]string
	GetName() string
}

type BaseEvent struct {
	Name       string
	Attributes map[string]string
}

type UserEvent struct {
	BaseEvent
}

type PropertyEvent struct {
	BaseEvent
}

type OtherEvent struct {
	BaseEvent
}

func (e *BaseEvent) GetName() string {
	return e.Name
}

func (e *BaseEvent) GetAttributes() map[string]string {
	return e.Attributes
}
