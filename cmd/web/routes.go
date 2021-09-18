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
	r.Use(SessionLoad)

	r.Get("/", handlers.Repo.Home)
	r.Get("/about", handlers.Repo.About)
	r.Get("/millhouse", handlers.Repo.MillHouse)
	r.Get("/rosehill", handlers.Repo.RoseHill)
	r.Get("/search-availability", handlers.Repo.Search)
	r.Get("/contact", handlers.Repo.Contact)


	// Creates a file server which serves files out of the "./ui/static" directory.
	fileServer := http.FileServer(http.Dir("./ui/static"))

	r.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return r
}
