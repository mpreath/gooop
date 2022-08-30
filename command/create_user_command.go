package command

import (
	"fmt"
)

type User struct {
	Name string
}

type CreateUserCommand struct {
	Command
}

func (c *CreateUserCommand) Execute() {
	objectType := fmt.Sprintf("%T", c.Object)
	if objectType == "command.User" {
		c.Result = c.Object
	} else {
		c.Result = nil
		c.Error = fmt.Errorf("execute: invalid object type")
	}
}
