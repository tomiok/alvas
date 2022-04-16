package useradmin

import (
	"github.com/tomiok/alvas/internal/views/login"
	"net/http"
)

func LoginViewHandler(w http.ResponseWriter, r *http.Request) {
	login.LoginViewHandler(w, nil)
}
