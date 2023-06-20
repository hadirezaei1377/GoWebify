package main

import (
	"GoWebify/pkg/config"
	"GoWebify/pkg/handlers"
	"net/http"

	"github.com/go-chi/chi" // router
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler { // return a route for application

	// middlewares setup
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)    // adds the recoverer middleware to the chi router
	mux.Use(NoSurf)                  // This adds custom middleware to the chi router that prevents CSRF
	mux.Use(SessionLoad)             // adds custom middleware to the chi router that loads the session data for each HTTP request
	mux.Get("/", handlers.Repo.Home) // Route definitions
	mux.Get("/about", handlers.Repo.About)

	return mux
}
