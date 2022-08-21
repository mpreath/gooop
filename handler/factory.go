package handler

import "fmt"

func New(handler_type string) (Handler, error) {
	switch handler_type {
	case "domain":
		return &DomainHandler{BaseHandler: BaseHandler{next: nil}}, nil
	case "log":
		return &LogHandler{BaseHandler: BaseHandler{next: nil}}, nil
	default:
		return nil, fmt.Errorf("unknown handler type \"%s\"", handler_type)
	}
}
