package database

import (
	"time"

	"github.com/spidernest-go/logger"
)

func (b *Board) New() error {
	b.DateCreated = time.Now()
	b.DateModified = time.Unix(0, 0)
	// TODO: Make sure something doesn't already exist in the spot [id, track_id]
	_, err := db.InsertInto("boards").
		Values(b).
		Exec()

	if err != nil {
		logger.Error().
			Err(err).
			Msg("Music Track could not be inserted into the table.")
	}

	return err
}
