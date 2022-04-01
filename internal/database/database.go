package database

import (
	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DBInstance *gorm.DB

func Init() {
	db, err := gorm.Open(sqlite.Open("alvas.db"), &gorm.Config{})

	if err != nil {
		log.Fatal().Err(err)
	}

	DBInstance = db
}
