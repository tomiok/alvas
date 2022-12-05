package repository

import (
	"github.com/tomiok/alvas/internal/customer"
	"github.com/tomiok/alvas/pkg"
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

func (r repository) Create(dto customer.CreateCustomer) (*customer.Customer, error) {
	pass, _ := pkg.HashPassword(dto.Password)
	c, err := customer.Create(dto.Name, dto.Address, dto.Email, pass)

	if err != nil {
		return nil, err
	}

	err = r.db.Create(&c).Error

	if err != nil {
		return nil, err
	}

	return c, err
}

func (r repository) Lookup(email string) (*customer.Customer, error) {
	var c customer.Customer
	err := r.db.Find(&c, "email=?", email).Error

	if err != nil {
		return nil, err
	}

	return &c, nil
}
