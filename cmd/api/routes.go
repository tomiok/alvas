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
	"os"
	"path/filepath"
	"strings"
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

func fileServer(r chi.Router) {
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "static"))
	fs(r, "/static", filesDir)

}

// fs conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func fs(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
