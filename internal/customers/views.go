package customers

import (
	"github.com/tomiok/alvas/internal/views/login"
	"github.com/tomiok/alvas/pkg/render"
	"net/http"
)

func NewCustomerViewHandler(w http.ResponseWriter, r *http.Request) {
	render.TemplateRender(w, "new_customer.page.tmpl", nil)
}

func LoginViewHandler(w http.ResponseWriter, r *http.Request) {
	login.LoginViewHandler(w, nil)
}
