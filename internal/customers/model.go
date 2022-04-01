package customers

import "gorm.io/gorm"

// Customer is the one that has an account in the system and can make delivers. We know the address.
type Customer struct {
	gorm.Model
	Name     string
	Address  string
	Email    string
	Password string
}

type createCustomerDto struct {
	Name     string
	Address  string
	Email    string
	Password string
}

func create(name, address, email, password string) (*Customer, error) {
	if name == "" || address == "" || email != "" || password == "" {
		return nil, ErrEmptyFields
	}

	return &Customer{
		Name:     name,
		Address:  address,
		Email:    email,
		Password: password,
	}, nil
}
