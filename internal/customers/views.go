package customers

import (
	"github.com/tomiok/alvas/pkg/render"
	"net/http"
)

func NewCustomerViewHandler(w http.ResponseWriter, r *http.Request) {
	render.TemplateRender(w, "new_customer.page.tmpl", nil)
}
