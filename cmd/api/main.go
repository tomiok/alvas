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
	db := database.New()

	//migrations
	err := migrate(db)

	if err != nil {
		log.Error().Err(err)
		log.Fatal().Msg("cannot migrate :(")
	}

	// template cache
	config.Init(render.TemplateRenderCache)

	deps := NewDependencies()
	r := routesSetup(deps)
	s := newServer("3333", r)

	s.start()
}

func migrate(db *gorm.DB) error {
	return migrations.Migrate(db)
}
