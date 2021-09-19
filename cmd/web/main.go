package main

import (
	"github.com/alexedwards/scs/v2"
	"github.com/tklara86/away/internal/config"
	"github.com/tklara86/away/internal/handlers"
	"github.com/tklara86/away/internal/render"
	"log"
	"net/http"
	"os"
	"time"
)

var app config.AppConfig
var session *scs.SessionManager


// main is the main application function
func main() {

	app.InProduction = false

	var app config.AppConfig

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session


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


	srv := http.Server{
		Addr: ":4000",
		Handler: mux,
	}

	logger.Printf("starting server on port:%s", srv.Addr)
	err = srv.ListenAndServe()
	log.Fatal(err)
}
