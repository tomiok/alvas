package useradmin

import (
	"github.com/tomiok/alvas/pkg/users"
	"gorm.io/gorm"
)

type Repository interface {
	CreateAdmin(dto createAdminDto) (*Admin, error)
	Lookup(email string) (*Admin, error)
}

type repo struct {
	db *gorm.DB
}

func newRepo(db *gorm.DB) *repo {
	return &repo{
		db: db,
	}
}

func (r repo) CreateAdmin(dto createAdminDto) (*Admin, error) {
	password, _ := users.HashPassword(dto.Password)
	admin := createAdmin(dto.Email, dto.Name, password)
	err := r.db.Create(admin).Error

	if err != nil {
		return nil, err
	}

	return admin, nil
}

func (r repo) Lookup(email string) (*Admin, error) {
	var admin Admin
	err := r.db.First(&admin, "email=?", email).Error

	if err != nil {
		return nil, err
	}

	return &admin, nil
}
