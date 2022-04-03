package migrations

import (
	"github.com/tomiok/alvas/internal/customers"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := customers.Migrate(db)

	if err != nil {
		return err
	}

	return nil
}
