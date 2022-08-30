package handler

type Event interface {
	GetAttributes() map[string]string
	SetAttributes(map[string]string)
	GetName() string
	SetName(string)
}

type BaseEvent struct {
	Name       string
	Attributes map[string]string
}

func (e *BaseEvent) GetName() string {
	return e.Name
}

func (e *BaseEvent) SetName(name string) {
	e.Name = name
}

func (e *BaseEvent) GetAttributes() map[string]string {
	return e.Attributes
}

func (e *BaseEvent) SetAttributes(attributes map[string]string) {
	e.Attributes = attributes
}

func NewEvent(name string, attributes map[string]string) Event {
	switch attributes["object"] {
	case "user":
		return &UserEvent{
			BaseEvent: BaseEvent{
				Name:       name,
				Attributes: attributes,
			},
		}
	case "property":
		return &PropertyEvent{
			BaseEvent: BaseEvent{
				Name:       name,
				Attributes: attributes,
			},
		}
	default:
		return &OtherEvent{
			BaseEvent: BaseEvent{
				Name:       name,
				Attributes: attributes,
			},
		}
	}
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
