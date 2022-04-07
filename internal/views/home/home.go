package home

import (
	"github.com/tomiok/alvas/pkg/render"
	"net/http"
)

func RenderView(w http.ResponseWriter, r *http.Request) {
	render.TemplateRender(w, "home.page.tmpl", nil)
}
