package main

import (
	"encoding/gob"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/tomiok/alvas/internal/customer"
	customerHandler "github.com/tomiok/alvas/internal/customer/handler"
	customerRepository "github.com/tomiok/alvas/internal/customer/repository"
	"github.com/tomiok/alvas/internal/database"
	deliveryHandler "github.com/tomiok/alvas/internal/delivery/handler"
	"github.com/tomiok/alvas/internal/user"
	userHandler "github.com/tomiok/alvas/internal/user/handler"
	userRepository "github.com/tomiok/alvas/internal/user/repository"
	"gorm.io/gorm"
)

type dependencies struct {
	db *gorm.DB

	userHandler     *userHandler.Handler
	customerHandler *customerHandler.Handler
	deliveryHandler *deliveryHandler.Handler

	session *scs.SessionManager
}

func NewDependencies() *dependencies {
	// registrations
	gob.Register(customer.SessCustomer{})

	db := database.New()

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

	_userHandler := userHandler.New(userService, session)
	_customerHandler := customerHandler.New(customerService, session)

	_deliveryHandler := deliveryHandler.New(session)
	return &dependencies{
		userHandler:     _userHandler,
		customerHandler: _customerHandler,
		deliveryHandler: _deliveryHandler,
		session:         session,
	}
}
