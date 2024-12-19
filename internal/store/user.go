package store

import "database/sql"

type UserStore struct {
	db *sql.DB
}

func (u *UserStore) Create() {
	// create user
}

func (u *UserStore) Update() {
	// update user
}

func (u *UserStore) Delete() {
	// delete user
}

func (u *UserStore) GetAllByToken() {
	// get user
}
