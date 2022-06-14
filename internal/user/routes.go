package user

import (
	"github.com/go-chi/chi/v5"
	"github.com/tomiok/alvas/internal/user/handler"
)

func Routes(web *handler.Handler) chi.Router {
	r := chi.NewRouter()
	r.Post("/", web.CreateAdminHandler)
	return r
}
