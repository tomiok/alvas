package handler

import (
	"github.com/gorilla/csrf"
	"github.com/tomiok/alvas/pkg/render"
	"net/http"
)

type Handler struct {
}

func New() *Handler {
	return &Handler{}
}

func (h Handler) SendPackageView() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		render.TemplateRender(w, "delivery.page.tmpl", &render.TemplateData{
			Data: map[string]interface{}{
				csrf.TemplateTag: csrf.TemplateField(r),
			},
		})
	}
}
