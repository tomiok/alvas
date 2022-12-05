package user

import (
	"errors"

	"github.com/tomiok/alvas/pkg"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

var ErrBadLogin = errors.New("cannot log in, please check your credentials")

type Admin struct {
	gorm.Model
	email    string
	name     string
	password string
}

func (Admin) TableName() string {
	return "admin_users"
}

type adminDto struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func (a Admin) toDto() *adminDto {
	return &adminDto{
		ID:    a.Model.ID,
		Email: a.email,
		Name:  a.name,
	}
}

func CreateAdmin(email, name, pass string) *Admin {
	return &Admin{
		email:    email,
		name:     name,
		password: pass,
	}
}

type Service interface {
	Create(email, name, pass string) (*adminDto, error)
	LogIn(email, password string) (*adminDto, error)
}

type Repository interface {
	CreateAdmin(email, name, pass string) (*Admin, error)
	Lookup(email string) (*Admin, error)
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

	if !pkg.DoPasswordsMatch(admin.password, password) {
		return nil, ErrBadLogin
	}

	res := admin.toDto()

	log.Info().Msg("admin logged OK")
	return res, nil
}

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&Admin{})

	return err
}
