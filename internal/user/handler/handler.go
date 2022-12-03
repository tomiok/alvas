package handler

import (
	"encoding/json"
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/tomiok/alvas/internal/user"
	"github.com/tomiok/alvas/pkg/web"
)

type Handler struct {
	s    user.Service
	sess *scs.SessionManager
}

func New(service user.Service, sess *scs.SessionManager) *Handler {
	return &Handler{
		s:    service,
		sess: sess,
	}
}

func (h Handler) CreateAdminHandler() func(w http.ResponseWriter, r *http.Request) {
	type createAdminDto struct {
		Email    string `json:"email"`
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		var req createAdminDto
		body := r.Body

		defer func() {
			_ = body.Close()
		}()

		err := json.NewDecoder(body).Decode(&req)

		if err != nil {
			web.ResponseBadRequest(w, "cannot decode request for create admin user", err)
			return
		}

		res, err := h.s.Create(req.Email, req.Name, req.Password)

		if err != nil {
			web.ResponseBadRequest(w, "cannot create admin user", err)
			return
		}

		web.Response2xx(w, http.StatusCreated, "admin user created", res)
	}
}

func (h Handler) LoginHandler() func(w http.ResponseWriter, r *http.Request) {
	type loginDto struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var dto loginDto
		body := r.Body

		defer func() {
			_ = body.Close()
		}()

		err := json.NewDecoder(body).Decode(&dto)

		if err != nil {
			web.ResponseBadRequest(w, "cannot decode", err)
			return
		}

		admin, err := h.s.LogIn(dto.Email, dto.Password)

		if err != nil {
			web.ResponseBadRequest(w, "cannot log in", err)
			return
		}

		web.Response2xx(w, http.StatusOK, "admin logged OK", admin)
	}
}
