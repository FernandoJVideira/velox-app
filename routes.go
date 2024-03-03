package main

import (
	"net/http"

	"github.com/FernandoJVideira/velox"
	"github.com/go-chi/chi/v5"
)

func (a *application) routes() *chi.Mux {
	// Middleware must be called before routes

	// Routes
	a.get("/", a.Handlers.Home)

	// Static Routes
	fileServer := http.FileServer(http.Dir("./public"))
	a.App.Routes.Handle("/public/*", http.StripPrefix("/public", fileServer))

	// Routes from velox
	a.App.Routes.Mount("/velox", velox.Routes())
	a.App.Routes.Mount("/api", a.ApiRoutes())

	return a.App.Routes
}
