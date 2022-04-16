package login

import (
	"github.com/tomiok/alvas/pkg/render"
	"net/http"
)

func LoginViewHandler(w http.ResponseWriter, data interface{}) {
	render.TemplateRender(w, "login.page.tmpl", &render.TemplateData{
		Data: data,
	})
}
