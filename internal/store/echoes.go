package store

import "database/sql"

type EchoesStore struct {
	db *sql.DB
}

func (e *EchoesStore) GetAll() {
	// get all echoes
}

func (e *EchoesStore) GetPopular() {
	// get popular echoes
}
func (e *EchoesStore) GetTrending() {
	// get trending echoes
}
