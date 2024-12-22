package main

import (
	"io"
	"net/http"
	"text/template"
	"time"

	"path/filepath"

	"github.com/codeshaine/echoes/internal/store"
)

const (
	staticDir   = "static"
	templateDir = "templates"
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

type Template struct {
	Template *template.Template
}

func (t *Template) Render(w io.Writer, name string, data any) {
	t.Template.ExecuteTemplate(w, name, data)
}

func NewTemplate() *Template {
	return &Template{
		Template: template.Must(template.ParseGlob(filepath.Join(templateDir, "*.html"))),
	}
}

func (a *App) mount() http.Handler {
	mux := http.NewServeMux()

	//static files for dev only
	fs := http.FileServer(http.Dir(staticDir))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fs))

	//for health check
	mux.HandleFunc("GET /health", a.CheckHealth)

	//home page
	mux.HandleFunc("GET /", a.RenderHomePage)
	mux.HandleFunc("GET /echoe/like/{id}", a.RenderEchoLikePage)

	// mux.HandleFunc("GET /user/echoe", a.GetEchoAllByToken)
	// mux.HandleFunc("POST /api/v1/user/echoe/create", a.CreateEcho)
	// mux.HandleFunc("PATCH /api/v1/user/echoe/update/:id", a.UpdateEcho)
	// mux.HandleFunc("DELETE /api/v1/user/echoe/delete/:id", a.DeleteEcho)
	// mux.HandleFunc("POST /api/v1/user/echoe/like/:id", a.LikeEcho)
	// mux.HandleFunc("DELETE /api/v1/user/echoe/unlike/:id", a.UnlikeEcho)

	// //render main page --- one for fetching all the echoes newest(default)
	// //filter one for fetching all the echoes olderst
	// //filter  popular/trending
	// //filter by most liked

	// mux.HandleFunc("GET /api/v1/echoes", a.GetEchoAll)
	// mux.HandleFunc("GET /api/v1/echoes/newest", a.GetEchoAllNewest)
	// mux.HandleFunc("GET /api/v1/echoes/oldest", a.GetEchoAllOldest)
	// mux.HandleFunc("GET /api/v1/echoes/popular", a.GetEchoAllPopular)

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
