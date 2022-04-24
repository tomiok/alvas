package customers

import (
	"encoding/json"
	"github.com/alexedwards/scs/v2"
	"github.com/tomiok/alvas/pkg/render"
	"github.com/tomiok/alvas/pkg/webutils"
	"net/http"
)

type Web struct {
	Service
	*scs.SessionManager
}

func newWeb(s Service, session *scs.SessionManager) *Web {
	return &Web{
		Service:        s,
		SessionManager: session,
	}
}

func (h Web) CreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if t := r.URL.Query().Get("type"); t == "form" {
			render.TemplateRender(w, "new_customer.page.tmpl", &render.TemplateData{})
			return
		}
	}

	var dto createCustomerDto
	body := r.Body
	defer func() {
		_ = body.Close()
	}()

	err := json.NewDecoder(body).Decode(&dto)

	if err != nil {
		webutils.ResponseBadRequest(w, "cannot decode create customer request", err)
		return
	}

	customer, err := h.Create(dto)

	if err != nil {
		webutils.ResponseBadRequest(w, "cannot create customer", err)
		return
	}

	webutils.Response2xx(w, http.StatusCreated, "customer created", customer.toDto())
}

type logInReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type logInRes struct {
	Message string `json:"message"`
}

func (h Web) LoginHandler(w http.ResponseWriter, r *http.Request) {
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
