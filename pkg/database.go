package pkg

import (
	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func New() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("alvas.db"), &gorm.Config{})

	if err != nil {
		log.Fatal().Err(err)
	}

	return db
}
