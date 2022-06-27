package main

import (
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
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
	db := database.Init()

	//migrations
	err := migrate(db)

	if err != nil {
		log.Error().Err(err)
		log.Fatal().Msg("cannot migrate :(")
	}

	// template cache
	config.Init(render.TemplateRenderCache)

	// session
	session := scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.Secure = false // true in prod
	session.Cookie.SameSite = http.SameSiteLaxMode

	r := routesSetup(db, session)
	s := newServer("3333", r)

	s.start()
}

func migrate(db *gorm.DB) error {
	return migrations.Migrate(db)
}
