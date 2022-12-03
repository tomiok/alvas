package main

import (
	"github.com/gorilla/csrf"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

type Server struct {
	http.Server
}

func newServer(port string, r chi.Router) *Server {
	return &Server{
		Server: http.Server{
			Addr:    ":" + port,
			Handler: r,
		},
	}
}

func (s *Server) start() {
	log.Info().Msgf("server staring in port %s", s.Addr)

	http.ListenAndServe(s.Addr, csrf.Protect([]byte("32-byte-long-auth-key"))(s.Handler))
}
