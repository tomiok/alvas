package useradmin

import "github.com/go-chi/chi/v5"

func Routes(web *Web) chi.Router {
	r := chi.NewRouter()

	r.Post("/", web.CreateAdminHandler)
	r.Post("/login", web.LoginHandler)

	// views
	r.Get("/views/login", LoginViewHandler)
	return r
}
