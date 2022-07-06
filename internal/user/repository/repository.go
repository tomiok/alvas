package repository

import (
	"github.com/tomiok/alvas/internal/user"
	"github.com/tomiok/alvas/pkg/users"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r repository) CreateAdmin(email, name, pass string) (*user.Admin, error) {
	password, _ := users.HashPassword(pass)
	admin := user.CreateAdmin(email, name, password)
	err := r.db.Create(admin).Error

	if err != nil {
		return nil, err
	}

	return admin, nil
}

func (r repository) Lookup(email string) (*user.Admin, error) {
	var admin user.Admin
	err := r.db.First(&admin, "email=?", email).Error

	if err != nil {
		return nil, err
	}

	return &admin, nil
}
