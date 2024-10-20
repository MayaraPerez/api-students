package main

import (
	"github.com/rs/zerolog/log"
	"github.com/MayaraPerez/api-students/api"
)

func main() {

  server := api.NewServer()
  server.Routes()

  if err := server.Start(); err != nil {
    log.Fatal().Err(err).Msgf("Failed to start server", err.Error())
  }
}

