package useradmin

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	email    string
	name     string
	password string
}

func (Admin) TableName() string {
	return "admin_users"
}

type createAdminDto struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type adminDto struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func (a Admin) toDto() *adminDto {
	return &adminDto{
		ID:    a.Model.ID,
		Email: a.email,
		Name:  a.name,
	}
}

func createAdmin(email, name, pass string) *Admin {
	return &Admin{
		email:    email,
		name:     name,
		password: pass,
	}
}
