package main

import (
	"flag"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/ppvan/guardian/app"
	"github.com/ppvan/guardian/data"
)

func main() {

	flag.Parse()
	dsn := flag.String("db", "guardian.db", "Database file")

	db, err := data.OpenDB(*dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	model := &data.PasswordModel{
		Connection: *db,
	}

	infoLog := log.New(os.Stdout, "", 0)
	errorLog := log.New(os.Stdin, "", 0)

	app := &app.Application{
		PasswordModel: model,
		Args:          flag.Args(),
		InfoLog:       infoLog,
		ErrorLog:      errorLog,
	}

	err = app.Run()
	if err != nil {
		app.ErrorLog.Fatal(err)
	}

}
