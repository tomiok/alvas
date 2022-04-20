package customers

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func CustomerRoutes(w *Web) chi.Router {
	r := chi.NewRouter()
	r.Post("/", w.CreateHandler)

	// views
	r.Get("/views/new_customer", NewCustomerViewHandler)

	return r
}

func CustomerLogin(w *Web) func(w http.ResponseWriter, r *http.Request) {
	return w.LoginHandler
}
