package customers

import (
	"github.com/alexedwards/scs/v2"
	"gorm.io/gorm"
)

func New(db *gorm.DB, session *scs.SessionManager) *Web {
	repo := newRepository(db)
	service := NewService(repo)
	return newWeb(service, session)
}
