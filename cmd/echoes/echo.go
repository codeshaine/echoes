package main

import (
	"log"
	"net/http"
)

func (a *App) RenderHomePage(w http.ResponseWriter, r *http.Request) {
	// filter := r.URL.Query().Get("filter")

	ctx := r.Context()
	data, err := a.Store.Echoes.GetAll(ctx)
	if err != nil {
		log.Printf("failed to get echoes: %v", err)
		http.Error(w, "failed to get echoes", http.StatusInternalServerError)
		return
	}

	t := NewTemplate()
	t.Render(w, "home", data)
}

func (a *App) RenderEchoLikePage(w http.ResponseWriter, r *http.Request) {
	// health check
}

func (a *App) RenderEchoUnlikePage(w http.ResponseWriter, r *http.Request) {
	// health check
}
