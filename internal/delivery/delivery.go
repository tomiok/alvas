package delivery

import "gorm.io/gorm"

//Delivery is the physical object that we want to send
type Delivery struct {
	gorm.Model
	SenderID    uint
	From        string
	Destination string
	Weight      float64
}

type Repository interface {
	SaveDelivery(d Delivery) (*Delivery, error)
}

type Service struct {
	r Repository
}

func NewDeliveryService(r Repository) *Service {
	return &Service{
		r: r,
	}
}

func (s *Service) Create(d Delivery) (*Delivery, error) {
	return s.r.SaveDelivery(d)
}

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&Delivery{})
	if err != nil {
		return err
	}

	return nil
}
