package command

import (
	"fmt"
)

type Dispatcher struct {
	Commands *Repository
}

func NewDispatcher(commands *Repository) *Dispatcher {
	return &Dispatcher{Commands: commands}
}

func (d *Dispatcher) Run(name string) error {
	command, err := d.Commands.GetCommand(name)
	if err != nil {
		return fmt.Errorf("cannot execute command <%s>, %w", name, ErrCommandNotFound)
	}

	err = command.Command("")
	if err != nil {
		return err
	}
	return nil
}
