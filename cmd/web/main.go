package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

// config defines configuration settings for our application.
type config struct {
	port int
	env string
}

// application hold dependencies for our HTTP handlers, helpers, middleware
type application struct {
	config config
	logger *log.Logger
}


func main() {

	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "Server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")

	flag.Parse()

	logger := log.New(os.Stdout, "", log.LstdFlags)

	app := &application{
		config: cfg,
		logger: logger,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", app.Home)
	mux.HandleFunc("/about", app.About)

	// Creates a file server which serves files out of the "./ui/static" directory.
	fileServer := http.FileServer(http.Dir("./ui/static"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	srv := http.Server{
		Addr: fmt.Sprintf(":%d", cfg.port),
		Handler: mux,
	}

	logger.Printf("starting %s server on port:%s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}
