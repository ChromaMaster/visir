package command

import "errors"

// TODO(fede): Tiene sentido mover esto aquí?
var (
	ErrCommandNotFound          = errors.New("command not found")
	ErrCommandAlreadyRegistered = errors.New("command already registered")
)
