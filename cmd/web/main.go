package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"GoWebify/pkg/config"
	"GoWebify/pkg/handlers"
	"GoWebify/pkg/render"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

// global variables
var app config.AppConfig        // stores the application configuration settings
var session *scs.SessionManager // manages user sessions

func main() {
	app.InProduction = false // program is not running in a production environment

	// set up the session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	// initialize the HTML template cache
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
