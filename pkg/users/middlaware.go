package users

import (
	"github.com/alexedwards/scs/v2"
	"net/http"
)

func LoadSession(sess *scs.SessionManager) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return sess.LoadAndSave(next)
	}
}
