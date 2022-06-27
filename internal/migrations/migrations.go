package migrations

import (
	"github.com/tomiok/alvas/internal/customer"
	"github.com/tomiok/alvas/internal/user"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := customer.Migrate(db); err != nil {
		return err
	}

	if err := user.Migrate(db); err != nil {
		return err
	}

	return nil
}
