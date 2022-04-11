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

type customerDto struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Email   string `json:"email"`
}

type createCustomerDto struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func create(name, address, email, password string) (*Customer, error) {
	if name == "" || address == "" || email == "" || password == "" {
		return nil, ErrEmptyFields
	}

	return &Customer{
		Name:     name,
		Address:  address,
		Email:    email,
		Password: password,
	}, nil
}

func (c Customer) toDto() customerDto {
	return customerDto{
		Name:    c.Name,
		Address: c.Address,
		Email:   c.Email,
	}
}
