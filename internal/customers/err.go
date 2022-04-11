package customers

import "errors"

var (
	ErrEmptyFields = errors.New("cannot create Customer, name, address, password and email are required")
)
