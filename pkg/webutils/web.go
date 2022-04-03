package webutils

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
)

type badRequestDto = struct {
	Message string `json:"message"`
}

type okDto = struct {
	Message string      `json:"message"`
	Value   interface{} `json:"value"`
}

func ResponseBadRequest(w http.ResponseWriter, message string) {
	log.Error().Msg("cannot decode entity")
	w.WriteHeader(http.StatusBadRequest)
	_ = json.NewEncoder(w).Encode(&badRequestDto{
		Message: message,
	})
}

func Response2xx(w http.ResponseWriter, status int, message string, any interface{}) {
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(&okDto{
		Message: message,
		Value:   any,
	})
}
