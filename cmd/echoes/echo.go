package main

import "net/http"

func (a *App) GetEchoAll(w http.ResponseWriter, r *http.Request) {
	// get echo
}

func (a *App) GetEchoAllNewest(w http.ResponseWriter, r *http.Request) {
	// create echo
}

func (a *App) GetEchoAllOldest(w http.ResponseWriter, r *http.Request) {
	// update echo
}

func (a *App) GetEchoAllPopular(w http.ResponseWriter, r *http.Request) {
	// delete echo
}
