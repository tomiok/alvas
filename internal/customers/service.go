package customers

type Service interface {
	Create(dto createCustomerDto) (*Customer, error)
	LogIn(email, password string) error
}

type service struct {
	r Repository
}

func NewService(r Repository) *service {
	return &service{
		r: r,
	}
}

func (s service) Create(dto createCustomerDto) (*Customer, error) {
	return s.r.Create(dto)
}

func (s service) LogIn(email, password string) error {
	return s.r.Lookup(email, password)
}
