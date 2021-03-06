package database

import (
	"github.com/spidernest-go/logger"
)

func Remove(id, tid uint64) error {
	err := db.Collection("boards").
		Find(id).
		Where("track_id = ", tid).
		Delete()
	if err != nil {
		logger.Error().
			Err(err).
			Msg("Entry did not exist or could not be deleted.")
	}
	return err
}
