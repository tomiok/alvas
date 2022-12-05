package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/alexedwards/scs/v2"
	"github.com/gorilla/csrf"
	"github.com/tomiok/alvas/internal/customer"
	"github.com/tomiok/alvas/internal/delivery"
	"github.com/tomiok/alvas/pkg"
)

type Handler struct {
	service *delivery.Service
	*scs.SessionManager
}

func New(svc *delivery.Service, scs *scs.SessionManager) *Handler {
	return &Handler{
		service:        svc,
		SessionManager: scs,
	}
}

func (h Handler) SendPackageView() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		c := h.Get(r.Context(), "customer").(customer.SessCustomer)
		pkg.TemplateRender(w, "delivery.page.tmpl", &pkg.TemplateData{
			Data: map[string]interface{}{
				csrf.TemplateTag: csrf.TemplateField(r),
				"customerID":     c.ID,
			},
		})
	}
}

func (h Handler) Generate() func(w http.ResponseWriter, r *http.Request) {
	type req struct {
		SenderID string `json:"senderID"`
		Sender   string `json:"sender"`
		AddrFrom string `json:"addrFrom"`
		AddrTo   string `json:"addrTo"`
		Weight   string `json:"weight"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		var req req
		body := r.Body
		defer func() {
			_ = body.Close()
		}()

		token := csrf.Token(r)
		w.Header().Set("X-CSRF-Token", token)

		err := json.NewDecoder(body).Decode(&req)

		if err != nil {
			pkg.ResponseBadRequest(w, "cannot decode", err)
			return
		}

		senderID, _ := strconv.ParseUint(req.SenderID, 10, 64)
		weight, _ := strconv.ParseFloat(req.Weight, 64)
		d, err := h.service.Create(delivery.Delivery{
			SenderID:    uint(senderID),
			From:        req.AddrFrom,
			Destination: req.AddrTo,
			Weight:      weight,
		})

		if err != nil {
			pkg.ResponseInternalError(w, "cannot create delivery", err)
			return
		}

		h.Put(r.Context(), "delivery", d)
		uri := fmt.Sprintf("/delivery?id=%d", d.ID)
		http.Redirect(w, r, uri, http.StatusSeeOther)
	}
}

func (h Handler) GetInformation() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		pkg.TemplateRender(w, "delivery-info.page.tmpl", &pkg.TemplateData{Data: map[string]interface{}{
			csrf.TemplateTag: csrf.TemplateField(r),
		}})
	}
}
