package main

import (
	"os"
	"test-1/internal/configs"
	"test-1/internal/server"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	config, configErr := configs.NewConfig()

	if configErr != nil {
		panic(configErr)
	}

	httpServer := server.CreateServer(log.Logger, config)
	httpServer.Run()
}
