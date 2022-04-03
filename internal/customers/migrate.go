package customers

import "gorm.io/gorm"

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&Customer{})
	if err != nil {
		return err
	}

	return nil
}
