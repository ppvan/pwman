package main

import (
	"flag"
	"log"

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

	app := &app.Application{
		PasswordModel: model,
		Args:          flag.Args(),
	}

	err = app.Run()
	if err != nil {
		log.Fatal(err)
	}

}
