package data

import "database/sql"

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
	return nil, nil
}

func (model *PasswordModel) Add(password Password) error {
	return nil
}

func (model *PasswordModel) Delete(id int) error {
	return nil
}
