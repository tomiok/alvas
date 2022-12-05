package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/alexedwards/scs/v2"

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

func ResponseInternalError(w http.ResponseWriter, message string, err error) {
	log.Error().Msg(err.Error())
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(&badRequestDto{
		Message: message,
	})
}

func ResponseUnauthorized(w http.ResponseWriter, message string) {
	log.Error().Msg("unauthorized")
	w.WriteHeader(http.StatusUnauthorized)
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

func getStackTrace() (uintptr, string, int, error) {
	pc, file, line, ok := runtime.Caller(2)

	if !ok {
		return 0, "", 0, errors.New("cannot get stack trace")
	}

	return pc, file, line, nil
}

// Trace is a helper for logs format. Prints information about the file, line and function calling.
// Expected result: log_trace_test.go -> TestTrace:17
func Trace() string {
	pc, path, line, err := getStackTrace()

	if err != nil {
		return ""
	}

	funcCall := runtime.FuncForPC(pc).Name()

	return fmt.Sprintf("%s -> %s:%d", getBase(path), getFuncName(getBase(funcCall)), line)
}

func getFuncName(funcBase string) string {
	funcName := strings.Split(funcBase, ".")
	l := len(funcName)

	if l == 1 {
		return funcName[0]
	}

	return funcName[1]
}

func getBase(path string) string {
	return filepath.Base(path)
}

func LoadSession(sess *scs.SessionManager) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return sess.LoadAndSave(next)
	}
}
