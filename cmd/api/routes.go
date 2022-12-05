package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/tomiok/alvas/internal/customer"
	"github.com/tomiok/alvas/pkg"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routesSetup(deps *dependencies) chi.Router {
	sess := deps.session
	r := chi.NewRouter()

	// middlewares
	r.Use(middleware.Recoverer)
	r.Use(pkg.LoadSession(sess))

	// file server
	fileServer(r)

	// login
	r.Get("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			pkg.TemplateRender(w, "login.page.tmpl", &pkg.TemplateData{
				IsLoginReq: true,
			})
		}

		if r.Method == http.MethodPost {
			log.Println("POST")
		}
	})

	// home
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		td := pkg.NewTemplateData()
		if sess.Exists(r.Context(), pkg.SessCustomerID) {
			td.CustomerName = sess.GetString(r.Context(), pkg.SessCustomerName)
			td.IsLogged = true
			c := sess.Get(r.Context(), "customer").(customer.SessCustomer)

			td.Data["customerID"] = c.ID
			td.Data["customerName"] = c.Name
			td.Data["customerAddress"] = c.Address
		}

		pkg.TemplateRender(w, "home.page.tmpl", td)
	})

	// customer
	_customerHandler := deps.customerHandler
	r.Route("/customers", func(r chi.Router) {
		r.Post("/", _customerHandler.CreateHandler())
		r.Get("/", _customerHandler.CreateHandlerView)
	})

	// user
	_userHandler := deps.userHandler
	r.Route("/admins", func(r chi.Router) {
		r.Post("/", _userHandler.CreateAdminHandler())
	})

	// delivery
	r.Route("/delivery", func(r chi.Router) {
		r.Get("/send-package", deps.deliveryHandler.SendPackageView())
		r.Get("/", deps.deliveryHandler.GetInformation())
		r.Post("/generate", deps.deliveryHandler.Generate())
	})

	// ping route
	r.Get("/ping", func(w http.ResponseWriter, req *http.Request) {
		_, _ = w.Write([]byte("pong"))
	})

	return r
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
