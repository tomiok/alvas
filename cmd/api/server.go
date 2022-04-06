package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
	"net/http"
)

type Server struct {
	s http.Server
}

func newServer(port string, r chi.Router) *Server {
	return &Server{
		s: http.Server{
			Addr:    ":" + port,
			Handler: r,
		},
	}
}

func (s *Server) start() {
	log.Info().Msgf("server staring in port %s", s.s.Addr)
	log.Fatal().Err(s.s.ListenAndServe())
}
