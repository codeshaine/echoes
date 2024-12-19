package store

import "database/sql"

type Storage struct {
	Echoes interface {
		GetAll()
		GetTrending()
		GetPopular()
	}
	User interface {
		Create()
		Delete()
		Update()
		GetAllByToken()
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Echoes: &EchoesStore{db},
		User:   &UserStore{db},
	}
}
