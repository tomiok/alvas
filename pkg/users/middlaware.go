package users

import (
	"net/http"

	"github.com/alexedwards/scs/v2"
)

func LoadSession(sess *scs.SessionManager) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return sess.LoadAndSave(next)
	}
}
