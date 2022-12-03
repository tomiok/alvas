package handler

import (
	"github.com/alexedwards/scs/v2"
	"github.com/gorilla/csrf"
	"github.com/tomiok/alvas/internal/customer"
	"github.com/tomiok/alvas/pkg/render"
	"net/http"
)

type Handler struct {
	*scs.SessionManager
}

func New(scs *scs.SessionManager) *Handler {
	return &Handler{scs}
}

func (h Handler) SendPackageView() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		c := h.Get(r.Context(), "customer").(customer.SessCustomer)
		render.TemplateRender(w, "delivery.page.tmpl", &render.TemplateData{
			Data: map[string]interface{}{
				csrf.TemplateTag: csrf.TemplateField(r),
				"customerID":     c.ID,
			},
		})
	}
}

func (h Handler) Generate() func(w http.ResponseWriter, r *http.Request) {
	type req struct {
	}
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
