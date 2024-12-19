package main

import (
	"log"

	"github.com/codeshaine/echoes/internal/db"
	"github.com/codeshaine/echoes/internal/env"
	"github.com/codeshaine/echoes/internal/store"
)

func main() {
	config := Config{
		addr: env.GetString("ADDR", ":3000"),
		dbConfig: DbCofig{
			addr:         env.GetString("DB_ADDR", "postgres://postgres:postgres@localhost:5432/echoes?sslmode=disable"),
			maxOpenConns: 30,
			maxIdleConns: 30,
			maxIdleTime:  "15m",
		},
	}
	db, err := db.New(
		config.dbConfig.addr,
		config.dbConfig.maxOpenConns,
		config.dbConfig.maxIdleConns,
		config.dbConfig.maxIdleTime,
	)
	if err != nil {
		log.Fatal("DbError:", err)
	}
	log.Println("Database connected")
	defer db.Close()

	store := store.NewStorage(db)
	app := App{
		Config: config,
		Store:  store,
	}

	mux := app.mount()
	log.Printf("Server started at %s", config.addr)
	log.Fatal(app.Run(mux))
}
