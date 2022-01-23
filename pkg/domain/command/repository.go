package command

import (
	"fmt"
)

type Repository struct {
	commands map[string]Command
}

type Command struct {
	Name    string
	Command func(string) error
}

func NewRepository() *Repository {
	var c Repository
	c.commands = make(map[string]Command)
	return &c
}

func (c Repository) findCommand(name string) (Command, error) {
	if command, ok := c.commands[name]; ok {
		return command, nil
	}

	return Command{}, fmt.Errorf("command <%s> not found: %w", name, ErrCommandNotFound)
}

func (c *Repository) RegisterCommand(command Command) error {
	_, err := c.findCommand(command.Name)
	if err == nil {
		return fmt.Errorf("command <%s> is already registered: %w", command.Name, ErrCommandAlreadyRegistered)
	}
	c.commands[command.Name] = command
	return nil
}

func (c Repository) GetCommand(name string) (Command, error) {
	command, err := c.findCommand(name)
	if err != nil {
		return Command{}, err
	}

	return command, nil
}

func (c Repository) UnregisterCommand(name string) error {
	_, err := c.findCommand(name)
	if err != nil {
		return err
	}
	delete(c.commands, "name")
	return nil
}
