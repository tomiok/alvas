package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/tomiok/alvas/pkg"
	"gorm.io/gorm"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	//database
	db := pkg.New()

	//migrations
	err := migrate(db)

	if err != nil {
		log.Error().Err(err)
		log.Fatal().Msg("cannot migrate :(")
	}

	// template cache
	pkg.Init(pkg.TemplateRenderCache)

	deps := NewDependencies()
	r := routesSetup(deps)
	s := newServer("3333", r)

	s.start()
}

func migrate(db *gorm.DB) error {
	return Migrate(db)
}
