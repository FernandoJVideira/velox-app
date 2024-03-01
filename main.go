package main

import (
	"myapp/data"
	"myapp/handlers"
	"myapp/middleware"

	"github.com/FernandoJVideira/velox"
)

type application struct {
	App        *velox.Velox
	Handlers   *handlers.Handlers
	Models     data.Models
	Middleware *middleware.Middleware
}

func main() {
	v := initApplication()
	v.App.ListenAndServe()
}
