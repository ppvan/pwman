package app

import (
	"fmt"

	"github.com/ppvan/guardian/data"
)

type Application struct {
	PasswordModel *data.PasswordModel
	Args          []string
}

func (app *Application) Run() error {

	if len(app.Args) == 0 {
		return fmt.Errorf("no command specified")
	}
	command, err := parseCommand(app.Args[0])
	if err != nil {
		return err
	}

	if err := app.excuteCommand(command); err != nil {
		return err
	}

	return nil
}

func (app *Application) ListPassword() {
	passwords, err := app.PasswordModel.List()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, password := range passwords {
		fmt.Printf("site: %s\n", password.Website)
		fmt.Printf("username: %s\n", password.Username)
	}
}

func (app *Application) AddPassword() {
	fmt.Println("Add Password")
}

func (app *Application) DeletePassword() {
	fmt.Println("Delete Password")
}

func (app *Application) UpdatePassword() {
	fmt.Println("Update Password")
}

func (app *Application) Help() {
	fmt.Println("Help")
}

func (app *Application) ResetMasterPassword() {
	fmt.Println("Reset Master Password")
}
