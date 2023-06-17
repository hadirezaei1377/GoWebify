package main

import (
	"GoWebify/pkg/config"
	"GoWebify/pkg/handlers"
	"net/http"

	"github.com/bmizerany/pat" // routing
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
