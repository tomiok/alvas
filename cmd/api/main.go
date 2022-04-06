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
	database.Init()

	//migrations
	err := migrate(database.DBInstance)

	if err != nil {
		log.Error().Err(err)
		log.Fatal().Msg("cannot migrate :(")
	}

	r := routesSetup(database.DBInstance)
	s := newServer("3333", r)

	s.start()
}

func migrate(db *gorm.DB) error {
	return migrations.Migrate(db)
}
