package customers

import "gorm.io/gorm"

type Repository interface {
	Create(dto createCustomerDto) (*Customer, error)
	Lookup(email, password string) error
}

type repository struct {
	db *gorm.DB
}

func (r repository) Create(dto createCustomerDto) (*Customer, error) {
	customer, err := create(dto.Name, dto.Address, dto.Email, dto.Password)

	if err != nil {
		return nil, err
	}

	err = r.db.Create(&customer).Error

	if err != nil {
		return nil, err
	}

	return customer, err
}
