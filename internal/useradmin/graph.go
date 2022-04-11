package useradmin

import (
	"github.com/alexedwards/scs/v2"
	"gorm.io/gorm"
)

func New(db *gorm.DB, sess *scs.SessionManager) *Web {
	repo := newRepo(db)
	svc := newService(repo)
	return &Web{
		s:    svc,
		sess: sess,
	}
}
