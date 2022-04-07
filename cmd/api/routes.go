package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/tomiok/alvas/internal/customers"
	"github.com/tomiok/alvas/internal/views/home"
	"gorm.io/gorm"
	"net/http"
)

func routesSetup(db *gorm.DB) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)

	customerRoutes(db, r)
	pingRoute(r)
	homeRoute(r)
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

func homeRoute(r chi.Router) {
	r.Get("/", home.RenderView)
}
