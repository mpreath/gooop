package handler

import "fmt"

func New(handlerType string) (Handler, error) {
	switch handlerType {
	case "domain":
		return NewDomainHandler(), nil
	case "log":
		return NewLogHandler(), nil
	default:
		return nil, fmt.Errorf("unknown handler type \"%s\"", handlerType)
	}
}
