package customers

import "errors"

var (
	ErrEmptyFields = errors.New("cannot create Customer, name and address are required")
)
