package command

import (
	"errors"
	"fmt"
)

var (
	ErrCommandNotFound          = errors.New("command not found")
	ErrCommandAlreadyRegistered = errors.New("command already registered")
)

type Repository interface {
	RegisterCommand(command Command) error
	GetCommand(name string) (Command, error)
	UnregisterCommand(name string) error
}

type CommandRepository struct {
	commands map[string]Command
}

type Command struct {
	Name    string
	Command func(string) error
}

func NewRepository() *CommandRepository {
	var c CommandRepository
	c.commands = make(map[string]Command)
	return &c
}

func (c CommandRepository) findCommand(name string) (Command, error) {
	if command, ok := c.commands[name]; ok {
		return command, nil
	}

	return Command{}, fmt.Errorf("command <%s> not found: %w", name, ErrCommandNotFound)
}

func (c *CommandRepository) RegisterCommand(command Command) error {
	_, err := c.findCommand(command.Name)
	if err == nil {
		return fmt.Errorf("command <%s> is already registered: %w", command.Name, ErrCommandAlreadyRegistered)
	}
	c.commands[command.Name] = command
	return nil
}

func (c CommandRepository) GetCommand(name string) (Command, error) {
	command, err := c.findCommand(name)
	if err != nil {
		return Command{}, err
	}

	return command, nil
}

func (c CommandRepository) UnregisterCommand(name string) error {
	_, err := c.findCommand(name)
	if err != nil {
		return err
	}
	delete(c.commands, "name")
	return nil
}
