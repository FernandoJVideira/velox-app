package handlers

import (
	"github.com/FernandoJVideira/velox"
	"myapp/data"
	"net/http"
)

// Handlers is the type for handlers, and gives access to Celeritas and models
type Handlers struct {
	App    *velox.Velox
	Models data.Models
}

// Home is the handler to render the home page
func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	err := h.render(w, r, "home", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println(err)
	}
}
