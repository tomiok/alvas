package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/tomiok/alvas/internal/database"
	"github.com/tomiok/alvas/internal/migrations"
	"gorm.io/gorm"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg("application starting")
	database.Init()

	//migrations
	err := migrate(database.DBInstance)

	if err != nil {
		log.Error().Err(err)
		log.Fatal().Msg("cannot migrate :(")
	}
}

func migrate(db *gorm.DB) error {
	return migrations.Migrate(db)
}
