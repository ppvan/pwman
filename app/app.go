package app

import "fmt"

type Application struct {
}

func (app *Application) ListPassword() {
	fmt.Println("List Password")
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
