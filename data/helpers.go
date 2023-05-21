package data

import (
	"database/sql"
	"errors"
	"os"
)

func OpenDB(path string) (*sql.DB, error) {
	if err := createDBFile(path); err != nil {
		return nil, err
	}

	conn, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(); err != nil {
		return nil, err
	}

	return conn, nil
}

func createDBFile(path string) error {

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		file, err := os.Create(path)
		if err != nil {
			return err
		}
		defer file.Close()
	}

	return nil
}

func createTable(conn *sql.DB) error {
	schema := `BEGIN TRANSACTION;
	CREATE TABLE IF NOT EXISTS "users" (
		"username"	TEXT NOT NULL DEFAULT 'guardian' UNIQUE,
		"password"	TEXT NOT NULL,
		"id"	INTEGER,
		PRIMARY KEY("id" AUTOINCREMENT)
	);
	CREATE TABLE IF NOT EXISTS "passwords" (
		"id"	INTEGER,
		"website"	TEXT NOT NULL,
		"password"	TEXT NOT NULL,
		"username"	TEXT NOT NULL,
		"created_at"	DATETIME,
		"updated_at"	DATETIME,
		"user_id"	INTEGER,
		FOREIGN KEY("user_id") REFERENCES "users"("id"),
		PRIMARY KEY("id" AUTOINCREMENT)
	);
	CREATE INDEX IF NOT EXISTS "user_idx" ON "passwords" (
		"user_id"
	);
	COMMIT;
	`

	_, err := conn.Exec(schema)
	if err != nil {
		return err
	}

	return nil
}
