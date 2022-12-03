package handler

import (
	"encoding/json"
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/gorilla/csrf"
	"github.com/tomiok/alvas/internal/customer"
	"github.com/tomiok/alvas/pkg/render"
	"github.com/tomiok/alvas/pkg/web"
)

type Handler struct {
	customer.Service
	*scs.SessionManager
}

func New(service customer.Service, session *scs.SessionManager) *Handler {
	return &Handler{
		Service:        service,
		SessionManager: session,
	}
}

func (h Handler) CreateHandlerView(w http.ResponseWriter, r *http.Request) {
	render.TemplateRender(w, "new.customer.page.tmpl", &render.TemplateData{
		Data: map[string]interface{}{
			csrf.TemplateTag: csrf.TemplateField(r),
		},
	})
}

func (h Handler) CreateHandler() func(w http.ResponseWriter, r *http.Request) {
	type createCustomerDto struct {
		Name     string `json:"name"`
		Address  string `json:"address"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var dto createCustomerDto
		token := csrf.Token(r)
		w.Header().Set("X-CSRF-Token", token)

		body := r.Body
		defer func() {
			_ = body.Close()
		}()

		err := json.NewDecoder(body).Decode(&dto)

		if err != nil {
			web.ResponseBadRequest(w, "cannot decode create customer request", err)
			return
		}

		_customer, err := h.Create(customer.CreateCustomer{
			Name:     dto.Name,
			Address:  dto.Address,
			Email:    dto.Email,
			Password: dto.Password,
		})

		if err != nil {
			web.ResponseBadRequest(w, "cannot create customer", err)
			return
		}

		h.Put(r.Context(), web.SessCustomerID, _customer.ID)
		h.Put(r.Context(), web.SessCustomerName, _customer.Name)
		h.Put(r.Context(), web.SessIsLogged, true)
		h.Put(r.Context(), "customer", customer.SessCustomer{
			ID:      _customer.ID,
			Name:    _customer.Name,
			Address: _customer.Address,
			Email:   _customer.Email,
		})

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func (h Handler) LoginHandler() func(w http.ResponseWriter, r *http.Request) {
	type logInReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		var dto logInReq

		body := r.Body
		defer func() {
			_ = body.Close()

		}()

		err := h.LogIn(dto.Email, dto.Password)
		if err != nil {
			web.ResponseBadRequest(w, "cannot log in", err)
			return
		}

		web.Response2xx(w, http.StatusOK, "logged ok", nil)
	}
}
