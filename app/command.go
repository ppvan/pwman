package app

import (
	"errors"
	"fmt"
)

type Command string

var ErrNoCommand = errors.New("no command specified")

type ErrUnKnownCommand struct {
	Command string
}

func (e *ErrUnKnownCommand) Error() string {
	return fmt.Sprintf("unknown command `%s`", e.Command)
}

const (
	Help    Command = "help"
	Reset   Command = "reset"
	Version Command = "version"
	List    Command = "list"
	Add     Command = "add"
	Delete  Command = "delete"
	Update  Command = "update"
)

var commands = []Command{
	Help,
	Reset,
	Version,
	List,
	Add,
	Delete,
	Update,
}

func parseCommand(str string) (Command, error) {
	for _, command := range commands {
		if command == Command(str) {
			return command, nil
		}
	}

	return "", &ErrUnKnownCommand{Command: str}
}
