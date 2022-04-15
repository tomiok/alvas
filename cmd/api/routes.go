package main

import (
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/tomiok/alvas/internal/customers"
	"github.com/tomiok/alvas/internal/useradmin"
	"github.com/tomiok/alvas/internal/views/home"
	sessmid "github.com/tomiok/alvas/pkg/users"
	"gorm.io/gorm"
	"net/http"
)

func routesSetup(db *gorm.DB, sess *scs.SessionManager) chi.Router {
	r := chi.NewRouter()

	// middlewares
	r.Use(middleware.Recoverer)
	//r.Use(csrfmid.NoSurf())
	r.Use(sessmid.LoadSession)

	// file server
	fileServer(r)

	// application routes
	customerRoutes(db, r, sess)
	adminRoutes(db, r, sess)
	pingRoute(r)
	homeRoute(r)

	return r
}

func customerRoutes(db *gorm.DB, r chi.Router, session *scs.SessionManager) {
	web := customers.New(db, session)
	customerRoutes := customers.CustomerRoutes(web)
	r.Mount("/customers", customerRoutes)
}

func adminRoutes(db *gorm.DB, r chi.Router, sess *scs.SessionManager) {
	web := useradmin.New(db, sess)
	adminRoutes := useradmin.Routes(web)
	r.Mount("/admins", adminRoutes)
}

func pingRoute(r chi.Router) {
	r.Get("/ping", func(w http.ResponseWriter, req *http.Request) {
		_, _ = w.Write([]byte("pong"))
	})
}

func homeRoute(r chi.Router) {
	r.Get("/", home.RenderView)
}

func fileServer(r *chi.Mux) {
	fs := http.FileServer(http.Dir("./static/"))
	r.Handle("/static/*", fs)
}
