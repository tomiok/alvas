package customers

import (
	"github.com/tomiok/alvas/pkg/users"
	"gorm.io/gorm"
)

type Repository interface {
	Create(dto createCustomerDto) (*Customer, error)
	Lookup(email string) (*Customer, error)
}

type repository struct {
	db *gorm.DB
}

func newRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r repository) Create(dto createCustomerDto) (*Customer, error) {
	pass, _ := users.HashPassword(dto.Password)
	customer, err := create(dto.Name, dto.Address, dto.Email, pass)

	if err != nil {
		return nil, err
	}

	err = r.db.Create(&customer).Error

	if err != nil {
		return nil, err
	}

	return customer, err
}

func (r repository) Lookup(email string) (*Customer, error) {
	var c Customer
	err := r.db.Find(&c, "email=?", email).Error

	if err != nil {
		return nil, err
	}

	return &c, nil
}
