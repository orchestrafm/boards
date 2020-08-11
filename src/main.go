package main

import (
	"github.com/orchestrafm/boards/src/database"
	"github.com/orchestrafm/boards/src/routing"
	"github.com/spidernest-go/logger"
)

func main() {
	err := database.Connect()
	if err != nil {

		logger.Error().
			Err(err).
			Msg("MySQL Database could not be attached to.")
	}
	routing.ListenAndServe()
}
