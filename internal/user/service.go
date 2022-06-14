package user

import (
	"errors"
	"github.com/rs/zerolog/log"
	"github.com/tomiok/alvas/pkg/users"
)

var ErrBadLogin = errors.New("cannot log in, please check your credentials")

type Service interface {
	Create(email, name, pass string) (*adminDto, error)
	LogIn(email, password string) (*adminDto, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) *service {
	return &service{
		repo: r,
	}
}

func (s service) Create(email, name, pass string) (*adminDto, error) {
	admin, err := s.repo.CreateAdmin(email, name, pass)

	if err != nil {
		return nil, err
	}

	res := admin.toDto()

	return res, nil
}

func (s service) LogIn(email, password string) (*adminDto, error) {
	admin, err := s.repo.Lookup(email)

	if err != nil {
		return nil, err
	}

	if !users.DoPasswordsMatch(admin.password, password) {
		return nil, ErrBadLogin
	}

	res := admin.toDto()

	log.Info().Msg("admin logged OK")
	return res, nil
}
