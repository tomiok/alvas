package useradmin

import "gorm.io/gorm"

type Repository interface {
	CreateAdmin(dto createAdminDto) (*Admin, error)
	Lookup(email, password string) (*Admin, error)
}

type repo struct {
	db *gorm.DB
}

func newRepo(db *gorm.DB) *repo {
	return &repo{
		db: db,
	}
}

func (r repo) CreateAdmin(dto createAdminDto) (*Admin, error) {
	admin := createAdmin(dto.Email, dto.Name, dto.Password)
	err := r.db.Create(admin).Error

	if err != nil {
		return nil, err
	}

	return admin, nil
}

func (r repo) Lookup(email, password string) (*Admin, error) {
	var admin Admin
	err := r.db.First(&admin, "email=?", email).Error

	if err != nil {
		return nil, err
	}

	return &admin, nil
}
