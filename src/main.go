package main

import (
	"os"
	"path/filepath"

	"github.com/orchestrafm/boards/src/database"
	"github.com/orchestrafm/boards/src/objstore"
	"github.com/orchestrafm/boards/src/routing"
	"github.com/spidernest-go/logger"
)

func main() {
	os.Mkdir(filepath.Join(os.TempDir(), "orchestrafm"), os.ModeDir)

	err := database.Connect()
	if err != nil {

		logger.Error().
			Err(err).
			Msg("MySQL Database could not be attached to.")
	}
	database.Synchronize()
	objstore.Login()
	routing.ListenAndServe()
}
