package useradmin

import (
	"encoding/json"
	"github.com/alexedwards/scs/v2"
	"github.com/tomiok/alvas/pkg/webutils"
	"net/http"
)

type Web struct {
	s    Service
	sess *scs.SessionManager
}

func (h Web) CreateAdminHandler(w http.ResponseWriter, r *http.Request) {
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

func (h Web) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var dto LoginDto
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
