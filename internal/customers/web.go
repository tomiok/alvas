package customers

import (
	"encoding/json"
	"github.com/tomiok/alvas/pkg/webutils"
	"net/http"
)

type Web struct {
	Service
}

func NewWeb(s Service) Web {
	return Web{Service: s}
}

func (h Web) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var dto createCustomerDto
	body := r.Body
	defer func() {
		_ = body.Close()
	}()

	err := json.NewDecoder(body).Decode(&dto)

	if err != nil {
		webutils.ResponseBadRequest(w, "cannot decode create customer request")
		return
	}

	customer, err := h.Create(dto)

	if err != nil {
		webutils.ResponseBadRequest(w, "cannot create customer")
		return
	}

	webutils.Response2xx(w, http.StatusCreated, "customer created", customer.toDto())
}
