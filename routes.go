package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *application) routes() *chi.Mux {
	// Middleware must be called before routes
	a.use(a.Middleware.CheckRemember)

	// Routes
	a.get("/", a.Handlers.Home)

	// Static Routes
	fileServer := http.FileServer(http.Dir("./public"))
	a.App.Routes.Handle("/public/*", http.StripPrefix("/public", fileServer))

	return a.App.Routes
}
