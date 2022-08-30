package command

type CommandInterface interface {
	Execute() (interface{}, error)
}

type Command struct {
	Result interface{}
	Error  error
	Object interface{}
}
