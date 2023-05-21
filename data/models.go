package data

import (
	"database/sql"
)

type Password struct {
	ID       int
	Username string
	Password string
	Website  string
	Notes    string
	Created  string
	Updated  string
}

type PasswordDAO interface {
	List() ([]Password, error)
	Add(password Password) error
	Delete(id int) error
	Update(password Password) error
}

type PasswordModel struct {
	Connection sql.DB
}

func (model *PasswordModel) List() ([]Password, error) {
	rows, err := model.Connection.Query("SELECT website, username, password FROM main.passwords")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var passwords []Password

	for rows.Next() {
		var password Password
		err := rows.Scan(&password.Website, &password.Username, &password.Password)
		if err != nil {
			return nil, err
		}
		passwords = append(passwords, password)
	}

	return passwords, nil
}

func (model *PasswordModel) Add(password Password) error {
	return nil
}

func (model *PasswordModel) Delete(id int) error {
	return nil
}
