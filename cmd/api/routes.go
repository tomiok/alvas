package main

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/csrf"
	customerHandler "github.com/tomiok/alvas/internal/customer/handler"
	userHandler "github.com/tomiok/alvas/internal/user/handler"
	"github.com/tomiok/alvas/pkg/render"
	sessmid "github.com/tomiok/alvas/pkg/users"
	"github.com/tomiok/alvas/pkg/webutils"
	"gorm.io/gorm"
)

func routesSetup(db *gorm.DB, sess *scs.SessionManager) chi.Router {
	r := chi.NewRouter()

	// middlewares
	r.Use(middleware.Recoverer)
	csrfMiddleware := csrf.Protect([]byte("32-byte-long-auth-key"))
	r.Use(csrfMiddleware)
	r.Use(sessmid.LoadSession(sess))

	// file server
	fileServer(r)

	// login
	r.Get("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			render.TemplateRender(w, r, "login.page.tmpl", &render.TemplateData{
				IsLoginReq: true,
			})
		}
	})

	// home
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		var td = &render.TemplateData{}
		if sess.Exists(r.Context(), webutils.SessCustomerID) {
			td.CustomerName = sess.GetString(r.Context(), webutils.SessCustomerName)
			td.IsLogged = true
		}

		render.TemplateRender(w, r, "home.page.tmpl", td)
	})

	// application routes
	_customerHandler := customerHandler.NewHandler(db, sess)
	r.Route("/customer", func(r chi.Router) {
		r.Post("/", _customerHandler.CreateHandler())
		r.Get("/", _customerHandler.CreateHandlerView)

	})

	_userHandler := userHandler.New(db, sess)
	r.Route("/admins", func(r chi.Router) {
		r.Post("/", _userHandler.CreateAdminHandler())
	})

	pingRoute(r)

	return r
}

func pingRoute(r chi.Router) {
	r.Get("/ping", func(w http.ResponseWriter, req *http.Request) {
		_, _ = w.Write([]byte("pong"))
	})
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
		panic("file server does not permit any URL parameters")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", http.StatusMovedPermanently).ServeHTTP)
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
