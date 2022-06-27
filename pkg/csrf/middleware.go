package csrfmid

import (
	"net/http"

	"github.com/gorilla/csrf"
)

func NoSurf() func(_ http.Handler) http.Handler {
	return csrf.Protect(
		[]byte("long-32-key-goes-here"),
		csrf.HttpOnly(false),
		csrf.Path("/"),
	)
}
