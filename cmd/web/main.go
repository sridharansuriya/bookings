package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/sridharansuriya/bookings/pkg/config"
	"github.com/sridharansuriya/bookings/pkg/handlers"
	"github.com/sridharansuriya/bookings/pkg/render"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main is the main application function
func main() {

	// Change this to true in production
	app.InProduction = false
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true //Persist session cookie after browser windows is closed
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("error creating template cache")
	}
	app.TemplateCache = tc
	app.UseCache = true

	render.SetAppConfig(&app)
	repo := handlers.SetAppConfig(&app)
	handlers.SetRepo(repo)

	// http.HandleFunc("/", handlers.Repo.Home)

	// http.HandleFunc("/about", handlers.Repo.About)

	log.Printf("Starting application on port: %s", portNumber)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	// http.ListenAndServe(portNumber, nil)
}
