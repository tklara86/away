package main

import (
	"github.com/tklara86/away/cmd/pkg/config"
	"github.com/tklara86/away/cmd/pkg/handlers"
	"github.com/tklara86/away/cmd/pkg/render"
	"log"
	"net/http"
	"os"
)




// main is the main application function
func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create tem[plate cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	logger := log.New(os.Stdout, "", log.LstdFlags)

	mux := Routes(&app)

	// Creates a file server which serves files out of the "./ui/static" directory.
	//fileServer := http.FileServer(http.Dir("./ui/static"))


	//mux.("/static/", http.StripPrefix("/static", fileServer))

	srv := http.Server{
		Addr: ":4000",
		Handler: mux,
	}

	logger.Printf("starting server on port:%s", srv.Addr)
	err = srv.ListenAndServe()
	log.Fatal(err)
}
