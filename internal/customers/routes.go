package customers

import (
	"github.com/go-chi/chi/v5"
)

func CustomerRoutes(w *Web) chi.Router {
	r := chi.NewRouter()
	r.Post("/", w.CreateHandler)

	// views
	r.Get("/views/new_customer", NewCustomerViewHandler)

	return r
}
