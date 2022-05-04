package useradmin

import "github.com/go-chi/chi/v5"

func Routes(web *Web) chi.Router {
	r := chi.NewRouter()
	r.Post("/", web.CreateAdminHandler)
	return r
}
