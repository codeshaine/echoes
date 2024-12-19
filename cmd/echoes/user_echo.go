package main

import "net/http"

func (a *App) GetEchoAllByToken(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetEchoAllByToken"))
	// get echo
}

func (a *App) CreateEcho(w http.ResponseWriter, r *http.Request) {
	// create echo
}

func (a *App) UpdateEcho(w http.ResponseWriter, r *http.Request) {
	// update echo
}

func (a *App) DeleteEcho(w http.ResponseWriter, r *http.Request) {
	// delete echo
}

func (a *App) LikeEcho(w http.ResponseWriter, r *http.Request) {
	// like echo
}

func (a *App) UnlikeEcho(w http.ResponseWriter, r *http.Request) {
	// unlike echo
}
