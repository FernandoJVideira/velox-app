package handlers

import (
	"myapp/data"
	"net/http"

	"github.com/FernandoJVideira/velox"
)

type Handlers struct {
	App    *velox.Velox
	Models data.Models
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	err := h.render(w, r, "home", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println(err)
	}
}
