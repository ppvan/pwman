package app

import (
	"fmt"
)

func (app *Application) excuteCommand(command Command) error {

	switch command {
	case Help:
		app.Help()
	case Reset:
		app.ResetMasterPassword()
	case Version:
		fmt.Println("Version")
	case List:
		app.ListPassword()
	case Add:
		app.AddPassword()
	case Delete:
		app.DeletePassword()
	case Update:
		app.UpdatePassword()
	default:
		// This should never happen
		panic("unknown command")
	}
	return nil
}
