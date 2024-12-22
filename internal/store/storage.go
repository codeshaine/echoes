package store

import (
	"context"
	"database/sql"
	"time"
)

const QueryTimeout = 5 * time.Second

type Storage struct {
	Echoes interface {
		GetAll(ctx context.Context) (EchoList, error)
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
