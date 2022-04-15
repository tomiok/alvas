package customers

import (
	"errors"
	"github.com/rs/zerolog/log"
	"github.com/tomiok/alvas/pkg/users"
)

var ErrBadLogin = errors.New("cannot log in, please check your credentials")

type Service interface {
	Create(dto createCustomerDto) (*Customer, error)
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

func (s service) Create(dto createCustomerDto) (*Customer, error) {
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
