package webutils

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

const (
	SessIsLogged     = "isLogged"
	SessCustomerID   = "customer_id"
	SessCustomerName = "customer_name"
)

type badRequestDto = struct {
	Message string `json:"message"`
}

type okDto = struct {
	Message string      `json:"message"`
	Value   interface{} `json:"value"`
}

func ResponseBadRequest(w http.ResponseWriter, message string, err error) {
	log.Error().Msg(err.Error())
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(&badRequestDto{
		Message: message,
	})
}

func Response2xx(w http.ResponseWriter, status int, message string, any interface{}) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(&okDto{
		Message: message,
		Value:   any,
	})
}
