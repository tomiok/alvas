package main

import (
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/tomiok/alvas/internal/customers"
	"github.com/tomiok/alvas/internal/views/home"
	csrfmid "github.com/tomiok/alvas/pkg/csrf"
	sessmid "github.com/tomiok/alvas/pkg/sess"
	"gorm.io/gorm"
	"net/http"
)

func routesSetup(db *gorm.DB, sess *scs.SessionManager) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(csrfmid.NoSurf())
	r.Use(sessmid.LoadSession)

	customerRoutes(db, r, sess)
	pingRoute(r)
	homeRoute(r)
	return r
}

func customerRoutes(db *gorm.DB, r chi.Router, session *scs.SessionManager) {
	web := customers.New(db, session)
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
