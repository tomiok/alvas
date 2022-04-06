package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/tomiok/alvas/internal/customers"
	"gorm.io/gorm"
	"net/http"
)

func routesSetup(db *gorm.DB) chi.Router {
	r := chi.NewRouter()
	customerRoutes(db, r)
	pingRoute(r)
	return r
}

func customerRoutes(db *gorm.DB, r chi.Router) {
	web := customers.New(db)
	customerR := customers.CustomerRoutes(web)

	r.Mount("/customers", customerR)
}

func pingRoute(r chi.Router) {
	r.Get("/ping", func(w http.ResponseWriter, req *http.Request) {
		_, _ = w.Write([]byte("pong"))
	})
}
