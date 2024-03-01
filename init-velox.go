package main

import (
	"log"
	"myapp/data"
	"myapp/handlers"
	"myapp/middleware"
	"os"

	"github.com/FernandoJVideira/velox"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Init Velox
	vel := &velox.Velox{}
	err = vel.New(path)
	if err != nil {
		log.Fatal(err)
	}

	vel.AppName = "myapp"

	myMiddleware := &middleware.Middleware{
		App: vel,
	}

	handlers := &handlers.Handlers{
		App: vel,
	}

	app := &application{
		App:        vel,
		Handlers:   handlers,
		Middleware: myMiddleware,
	}

	app.App.Routes = app.routes()
	app.Models = data.New(app.App.DB.Pool)
	handlers.Models = app.Models
	app.Middleware.Models = app.Models

	return app
}
