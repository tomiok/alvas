package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tomiok/alvas/internal/customer/repository"

	"github.com/gorilla/csrf"
	"github.com/tomiok/alvas/internal/customer"
	"gorm.io/gorm"

	"github.com/alexedwards/scs/v2"
	"github.com/tomiok/alvas/pkg/render"
	"github.com/tomiok/alvas/pkg/webutils"
)

type Handler struct {
	customer.Service
	*scs.SessionManager
}

func NewHandler(db *gorm.DB, session *scs.SessionManager) *Handler {
	repo := repository.NewRepository(db)
	service := customer.NewService(repo)
	return &Handler{
		Service:        service,
		SessionManager: session,
	}
}

func (h Handler) CreateHandlerView(w http.ResponseWriter, r *http.Request) {
	render.TemplateRender(w, r, "new.customer.page.tmpl", &render.TemplateData{})
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
			webutils.ResponseBadRequest(w, "cannot decode create customer request", err)
			return
		}

		_customer, err := h.Create(customer.CreateCustomer{
			Name:     dto.Name,
			Address:  dto.Address,
			Email:    dto.Email,
			Password: dto.Password,
		})

		if err != nil {
			webutils.ResponseBadRequest(w, "cannot create customer", err)
			return
		}
		h.Put(r.Context(), webutils.SessCustomerID, _customer.ID)
		h.Put(r.Context(), webutils.SessCustomerName, _customer.Name)
		h.Put(r.Context(), webutils.SessIsLogged, true)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

type logInReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h Handler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var dto logInReq

	body := r.Body
	defer func() {
		_ = body.Close()

	}()

	err := h.LogIn(dto.Email, dto.Password)
	if err != nil {
		webutils.ResponseBadRequest(w, "cannot log in", err)
		return
	}

	webutils.Response2xx(w, http.StatusOK, "logged ok", nil)
}
