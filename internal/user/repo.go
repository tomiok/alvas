package user

import (
	"github.com/tomiok/alvas/pkg/users"
	"gorm.io/gorm"
)

type Repository interface {
	CreateAdmin(email, name, pass string) (*Admin, error)
	Lookup(email string) (*Admin, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r repository) CreateAdmin(email, name, pass string) (*Admin, error) {
	password, _ := users.HashPassword(pass)
	admin := createAdmin(email, name, password)
	err := r.db.Create(admin).Error

	if err != nil {
		return nil, err
	}

	return admin, nil
}

func (r repository) Lookup(email string) (*Admin, error) {
	var admin Admin
	err := r.db.First(&admin, "email=?", email).Error

	if err != nil {
		return nil, err
	}

	return &admin, nil
}
