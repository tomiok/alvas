package main

import (
	"github.com/rs/zerolog"
	"github.com/tomiok/alvas/internal/database"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	database.Init()
}
