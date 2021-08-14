package main

import (
	"github.com/go-chi/chi"
	"github.com/tklara86/away/cmd/pkg/config"
	"github.com/tklara86/away/cmd/pkg/handlers"
	"net/http"
)

func Routes(app *config.AppConfig) http.Handler {
	r := chi.NewRouter()

	r.Use(NoSurf)

	r.Get("/", handlers.Repo.Home)
	r.Get("/about", handlers.Repo.About)

	return r
}
