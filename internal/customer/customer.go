package customer

import (
	"errors"

	"github.com/rs/zerolog/log"
	"github.com/tomiok/alvas/pkg/users"
	"gorm.io/gorm"
)

var ErrEmptyFields = errors.New("cannot create Customer, name, address, password and email are required")

// Customer is the one that has an account in the system and can make delivers. We know the address.
type Customer struct {
	gorm.Model
	Name     string
	Address  string
	Email    string
	Password string
}

type SessCustomer struct {
	ID      uint
	Name    string
	Address string
	Email   string
}

type CreateCustomer struct {
	Name     string
	Address  string
	Email    string
	Password string
}

type Repository interface {
	Create(dto CreateCustomer) (*Customer, error)
	Lookup(email string) (*Customer, error)
}

func Create(name, address, email, password string) (*Customer, error) {
	if name == "" || address == "" || email == "" || password == "" {
		return nil, ErrEmptyFields
	}

	return &Customer{
		Name:     name,
		Address:  address,
		Email:    email,
		Password: password,
	}, nil
}

var ErrBadLogin = errors.New("cannot log in, please check your credentials")

type Service interface {
	Create(dto CreateCustomer) (*Customer, error)
	LogIn(email, password string) error
}

type service struct {
	r Repository
}

func NewService(r Repository) *service {
	return &service{
		r: r,
	}
}

func (s service) Create(dto CreateCustomer) (*Customer, error) {
	return s.r.Create(dto)
}

func (s service) LogIn(email, password string) error {
	c, err := s.r.Lookup(email)

	if err != nil {
		return err
	}

	if !users.DoPasswordsMatch(c.Password, password) {
		return ErrBadLogin
	}

	log.Info().Msg("user logged in OK")
	return nil
}
