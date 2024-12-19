package main

import (
	"net/http"
	"time"

	"github.com/codeshaine/echoes/internal/store"
)

type App struct {
	Config Config
	Store  store.Storage
}

type Config struct {
	addr     string
	dbConfig DbCofig
}

type DbCofig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

func (a *App) mount() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", a.CheckHealth)

	//create a new echo
	//update an echo
	//delete an echo
	//like an echo
	//unlike an echo
	mux.HandleFunc("GET /api/v1/user/echoe", a.GetEchoAllByToken)
	mux.HandleFunc("POST /api/v1/user/echoe/create", a.CreateEcho)
	mux.HandleFunc("PATCH /api/v1/user/echoe/update/:id", a.UpdateEcho)
	mux.HandleFunc("DELETE /api/v1/user/echoe/delete/:id", a.DeleteEcho)
	mux.HandleFunc("POST /api/v1/user/echoe/like/:id", a.LikeEcho)
	mux.HandleFunc("DELETE /api/v1/user/echoe/unlike/:id", a.UnlikeEcho)

	//render main page --- one for fetching all the echoes newest(default)
	//filter one for fetching all the echoes olderst
	//filter  popular/trending
	//filter by most liked

	mux.HandleFunc("GET /api/v1/echoes", a.GetEchoAll)
	mux.HandleFunc("GET /api/v1/echoes/newest", a.GetEchoAllNewest)
	mux.HandleFunc("GET /api/v1/echoes/oldest", a.GetEchoAllOldest)
	mux.HandleFunc("GET /api/v1/echoes/popular", a.GetEchoAllPopular)

	return Logger(mux)
}

func (a *App) Run(mux http.Handler) error {
	server := &http.Server{
		Addr:         a.Config.addr,
		Handler:      mux,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  10 * time.Second,
	}
	return server.ListenAndServe()
}
