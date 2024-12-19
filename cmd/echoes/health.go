package main

import "net/http"

func (a *App) CheckHealth(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ECHOES IS ALIVE!"))
}
