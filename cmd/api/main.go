package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/tomiok/alvas/internal/database"
	"github.com/tomiok/alvas/internal/migrations"
	"github.com/tomiok/alvas/pkg/config"
	"github.com/tomiok/alvas/pkg/render"
	"gorm.io/gorm"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	//database
	database.Init()

	//migrations
	err := migrate(database.DBInstance)

	if err != nil {
		log.Error().Err(err)
		log.Fatal().Msg("cannot migrate :(")
	}

	// template cache
	config.Init(render.TemplateRenderCache)

	r := routesSetup(database.DBInstance)
	s := newServer("3333", r)

	s.start()
}

func migrate(db *gorm.DB) error {
	return migrations.Migrate(db)
}
