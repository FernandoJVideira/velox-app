package middleware

import (
	"myapp/data"

	"github.com/FernandoJVideira/velox"
)

type Middleware struct {
	App    *velox.Velox
	Models data.Models
}
