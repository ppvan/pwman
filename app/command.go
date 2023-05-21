package app

import "fmt"

type Command string

const UnKnownCommandError = "unknown command `%s`"

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

	return "", fmt.Errorf(UnKnownCommandError, str)
}
