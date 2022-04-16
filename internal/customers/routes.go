package customers

import "github.com/go-chi/chi/v5"

func CustomerRoutes(w Web) chi.Router {
	r := chi.NewRouter()
	r.Post("/", w.CreateHandler)
	r.Post("/login", w.LoginHandler)

	// views
	r.Get("/views/login", LoginViewHandler)
	r.Get("/views/new_customer", NewCustomerViewHandler)

	return r
}
