package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/ppvan/guardian/app"
)

func main() {

	app := &app.Application{}

	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		log.Fatal("You must specify a command")
	}

	switch args[0] {
	case "help":
		app.Help()
	case "reset":
		app.ResetMasterPassword()
	case "version":
		fmt.Println("Version")
	case "list":
		app.ListPassword()
	case "add":
		app.AddPassword()
	case "delete":
		app.DeletePassword()
	case "update":
		app.UpdatePassword()
	default:
		log.Fatalf("Unknown command %s", args[0])
	}
}
