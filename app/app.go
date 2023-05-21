package app

import (
	"log"

	"github.com/ppvan/guardian/data"
)

type Application struct {
	PasswordModel *data.PasswordModel
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	Args          []string
}

func (app *Application) Run() error {

	if len(app.Args) == 0 {
		return ErrNoCommand
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

func (app *Application) ListPassword() error {
	passwords, err := app.PasswordModel.List()
	if err != nil {
		return err
	}

	for _, password := range passwords {
		app.InfoLog.Printf("site: %s\n", password.Website)
		app.InfoLog.Printf("username: %s\n", password.Username)
	}

	return nil
}

func (app *Application) AddPassword() error {
	app.InfoLog.Println("Add Password")
	return nil
}

func (app *Application) DeletePassword() error {
	app.InfoLog.Println("Delete Password")
	return nil
}

func (app *Application) UpdatePassword() error {
	app.InfoLog.Println("Update Password")
	return nil
}

func (app *Application) Help() error {
	app.InfoLog.Println("Help")
	return nil
}

func (app *Application) ResetMasterPassword() error {
	app.InfoLog.Println("Reset Master Password")
	return nil
}
