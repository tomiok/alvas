package repository

import (
	"github.com/tomiok/alvas/internal/delivery"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewDeliveryRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) SaveDelivery(d delivery.Delivery) (*delivery.Delivery, error) {
	err := r.db.Create(&d).Error

	if err != nil {
		return nil, err
	}

	return &d, nil
}
