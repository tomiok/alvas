package handler

import (
	"encoding/json"
	"github.com/alexedwards/scs/v2"
	"github.com/tomiok/alvas/internal/user"
	"github.com/tomiok/alvas/pkg/webutils"
	"gorm.io/gorm"
	"net/http"
)

type Handler struct {
	s    user.Service
	sess *scs.SessionManager
}

func New(db *gorm.DB, sess *scs.SessionManager) *Handler {
	repo := user.NewRepository(db)
	svc := user.NewService(repo)
	return &Handler{
		s:    svc,
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
			webutils.ResponseBadRequest(w, "cannot decode request for create admin user", err)
			return
		}

		res, err := h.s.Create(req)

		if err != nil {
			webutils.ResponseBadRequest(w, "cannot create admin user", err)
			return
		}

		webutils.Response2xx(w, http.StatusCreated, "admin user created", res)
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
			webutils.ResponseBadRequest(w, "cannot decode", err)
			return
		}

		admin, err := h.s.LogIn(dto.Email, dto.Password)

		if err != nil {
			webutils.ResponseBadRequest(w, "cannot log in", err)
			return
		}

		webutils.Response2xx(w, http.StatusOK, "admin logged OK", admin)
	}
}
