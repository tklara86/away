package main

import (
	"encoding/gob"
	"fmt"
	"github.com/alexedwards/scs/v2"

	"github.com/tklara86/away/internal/config"
	"github.com/tklara86/away/internal/driver"
	"github.com/tklara86/away/internal/handlers"
	"github.com/tklara86/away/internal/helpers"
	"github.com/tklara86/away/internal/models"
	"github.com/tklara86/away/internal/render"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger


// main is the main application function
func main() {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")


	// what to store in the session
	gob.Register(models.Reservation{})
	gob.Register(models.User{})
	gob.Register(models.Room{})
	gob.Register(models.Restriction{})

	fmt.Println(os.Getenv("DB_HOST"))

	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.LstdFlags)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.LstdFlags|log.Lshortfile)
	app.ErrorLog = errorLog

	var app config.AppConfig

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	// connect to database
	log.Println("connecting to database")


	dsn := fmt.Sprintf("host=%s port=5432 dbname=%s user=%s password=%s\n", dbHost, dbName, dbUser,
		dbPassword)

	db, err := driver.ConnectSQL(dsn)

	if err != nil {
		log.Fatal("cannot connect to the database")
	}
	defer db.SQL.Close()

	log.Println("connected to the database")

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false


	repo := handlers.NewRepo(&app, db)

	handlers.NewHandlers(repo)

	render.NewRenderer(&app)

	helpers.NewHelpers(&app)

	mux := Routes(&app)


	srv := http.Server{
		Addr: ":4000",
		Handler: mux,
	}

	//logger.Printf("starting server on port:%s", srv.Addr)
	err = srv.ListenAndServe()
	log.Fatal(err)
}
