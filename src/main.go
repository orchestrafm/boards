package main

import (
	"github.com/orchestrafm/boards/src/database"
	"github.com/orchestrafm/boards/src/routing"
	"github.com/spidernest-go/logger"
)

func main() {
	err := database.Connect()
	logger.Error().
		Err(err).
		Msg("MySQL Database could not be attached to.")
	database.Synchronize()

	routing.ListenAndServe()
}
