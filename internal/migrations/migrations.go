package migrations

import (
	"github.com/tomiok/alvas/internal/customers"
	"github.com/tomiok/alvas/internal/user"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err1 := customers.Migrate(db)

	if err1 != nil {
		return err1
	}

	err2 := user.Migrate(db)

	if err2 != nil {
		return err2
	}

	return nil
}
