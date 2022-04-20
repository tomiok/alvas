package login

import (
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
	"github.com/tomiok/alvas/pkg/render"
	"net/http"
)

func MainLoginViewRoutes(r chi.Router) {
	r.Get("/login", mainLoginViewHandler)
	r.Get("/main_login", mainLoginHandler)
}

func mainLoginViewHandler(w http.ResponseWriter, r *http.Request) {
	render.TemplateRender(w, "login.page.tmpl", nil)
}

func mainLoginHandler(w http.ResponseWriter, r *http.Request) {
	value := r.FormValue("userType")
	log.Info().Msgf("value: %s", value)
}
