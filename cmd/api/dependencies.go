package main

import (
	"encoding/gob"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/csrf"
	"github.com/rs/zerolog/log"
	"github.com/tomiok/alvas/internal/delivery"
	deliveryRepository "github.com/tomiok/alvas/internal/delivery/repository"
	"github.com/tomiok/alvas/pkg"

	"github.com/alexedwards/scs/v2"
	"github.com/tomiok/alvas/internal/customer"
	customerHandler "github.com/tomiok/alvas/internal/customer/handler"
	customerRepository "github.com/tomiok/alvas/internal/customer/repository"
	deliveryHandler "github.com/tomiok/alvas/internal/delivery/handler"
	"github.com/tomiok/alvas/internal/user"
	userHandler "github.com/tomiok/alvas/internal/user/handler"
	userRepository "github.com/tomiok/alvas/internal/user/repository"
)

type dependencies struct {
	userHandler     *userHandler.Handler
	customerHandler *customerHandler.Handler
	deliveryHandler *deliveryHandler.Handler

	session *scs.SessionManager
}

func NewDependencies() *dependencies {
	// registrations
	gob.Register(customer.SessCustomer{})
	gob.Register(delivery.Delivery{})

	db := pkg.New()

	// session
	session := scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.Secure = false // true in prod
	session.Cookie.SameSite = http.SameSiteLaxMode

	_userRepository := userRepository.NewRepository(db)
	userService := user.NewService(_userRepository)

	_customerRepository := customerRepository.NewRepository(db)
	customerService := customer.NewService(_customerRepository)

	_deliveryRepository := deliveryRepository.NewDeliveryRepository(db)
	deliveryService := delivery.NewDeliveryService(_deliveryRepository)

	_userHandler := userHandler.New(userService, session)
	_customerHandler := customerHandler.New(customerService, session)
	_deliveryHandler := deliveryHandler.New(deliveryService, session)
	return &dependencies{
		userHandler:     _userHandler,
		customerHandler: _customerHandler,
		deliveryHandler: _deliveryHandler,
		session:         session,
	}
}

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

	log.Fatal().Err(http.ListenAndServe(s.Addr, csrf.Protect([]byte("32-byte-long-auth-key"))(s.Handler)))
}
