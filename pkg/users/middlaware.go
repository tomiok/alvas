package users

import (
	"github.com/alexedwards/scs/v2"
	"net/http"
)

func LoadSession(next http.Handler) http.Handler {
	return scs.New().LoadAndSave(next)
}
