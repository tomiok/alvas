package customers

import "gorm.io/gorm"

func New(db *gorm.DB) Web {
	repo := newRepository(db)
	service := NewService(repo)
	return NewWeb(service)
}
