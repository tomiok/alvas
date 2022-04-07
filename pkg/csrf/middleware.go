package csrfmid

import (
	"github.com/gorilla/csrf"
	"net/http"
)

func GetCSRFMid() func(_ http.Handler) http.Handler {
	return csrf.Protect([]byte("long-32-key-goes-here"))
}
