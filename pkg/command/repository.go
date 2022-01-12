package command

import (
	"errors"
	"fmt"
)

type Repository interface {
	Register(command Command) error
	Get(name string) (Command, error)
}

type CommandRepository struct {
	commands map[string]Command
}

type Command struct {
	Name    string
	Command func(string)
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

	return Command{}, errors.New(fmt.Sprintf("Command <%s> not found", name))
}

func (c *CommandRepository) Register(command Command) error {
	_, err := c.findCommand(command.Name)
	if err == nil {
		return errors.New(fmt.Sprintf("Command <%s> is already registered", command.Name))
	}
	c.commands[command.Name] = command
	return nil
}

func (c CommandRepository) Get(name string) (Command, error) {
	command, err := c.findCommand(name)
	if err != nil {
		return Command{}, err
	}

	return command, nil
}
